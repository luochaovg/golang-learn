package router

import (
	"context"
	"crypto/sha256"
	"dashboard-backend/pkg/gredis"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// InitScreenRouter 数据治理大屏监控 （demo）
func InitScreenRouter(r *gin.Engine) *gin.Engine {

	screenApi := r.Group("api/v2/monitor/screen/")

	var ctx = context.Background()
	redisClient, _ := gredis.GetRedisClient(gredis.MonitorScreenRedis)

	// Cache to store the last sent data hash
	var lastSentDataHash string
	var mu sync.Mutex

	screenApi.GET("/events", func(c *gin.Context) {
		// 设置响应头
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		// 订阅 Redis 通道
		pubsub := redisClient.Client.Subscribe(ctx, "game_metrics_channel")
		defer pubsub.Close()

		// Q: 前端页面刷新， redis 订阅通道的数量 就会增加 ， 订阅数据的增加会有影响吗
		// 是的，每次前端页面刷新都会导致一个新的订阅通道创建，这会导致 Redis 订阅通道数量增加，从而增加系统负载。
		// 虽然 Redis 可以处理大量的订阅，但在高流量情况下，过多的订阅会影响性能。
		// 你可以采取以下几种方法来优化这个问题：
		// 限制订阅的生命周期：确保订阅在客户端断开连接时终止。

		// 创建一个通道来监听客户端断开连接
		notify := c.Writer.CloseNotify()
		done := make(chan struct{})
		go func() {
			select {
			case <-notify:
				log.Println("Client disconnected")
				close(done)
			}
		}()

		// Q: SSE 默认支持了断线重连， 还需要启动心跳检测机制吗？
		// 虽然 SSE（Server-Sent Events）默认支持断线重连，但在某些情况下，心跳检测机制仍然是有用的。以下是一些需要心跳检测机制的原因和场景：
		//检测连接状态：
		//SSE 的断线重连机制依赖于浏览器检测到网络断开，但有时网络连接可能处于“假连接”状态（例如，连接未完全断开但无法通信）。
		//使用心跳检测可以更及时地发现这种情况并采取措施。
		//服务器主动检测客户端状态：
		//服务器可以利用心跳检测来及时了解客户端是否还在监听，从而进行资源管理（如清理失效的连接）。
		//防止长时间无活动：
		//即使 SSE 支持自动重连，长时间没有任何数据传输的连接可能会被一些中间设备（如防火墙或代理服务器）关闭。
		//定期发送心跳消息可以保持连接的活跃状态，防止被中间设备关闭。

		// 启动心跳检测
		go func() {
			ticker := time.NewTicker(30 * time.Second) // 每 30 秒发送一次心跳
			defer ticker.Stop()
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					c.SSEvent("heartbeat", "ping")
					c.Writer.Flush()
				}
			}
		}()

		// 循环发送消息给客户端
		for {
			select {
			case <-done:
				return
			default:
				var maxRetries = 5
				var retryInterval = time.Second * 3

				var msg *redis.Message
				var err error

				for retries := 0; retries < maxRetries; retries++ {
					msg, err = pubsub.ReceiveMessage(ctx)
					if err == nil {
						break
					}

					log.Printf("Error receiving message, retrying in %v... (%d/%d)\n", retryInterval, retries+1, maxRetries)
					time.Sleep(retryInterval)

					// 指数退避，增加重试间隔
					retryInterval *= 2
				}

				if err != nil {
					log.Println("Error receiving message after retries:", err)
					return
				}

				// 防抖逻辑：检查新数据是否与缓存的数据哈希相同
				mu.Lock()
				dataHash := fmt.Sprintf("%x", sha256.Sum256([]byte(msg.Payload)))
				if dataHash != lastSentDataHash {
					lastSentDataHash = dataHash
					c.SSEvent("message", msg.Payload)
					c.Writer.Flush()
				}
				mu.Unlock()
			}
		}
	})

	// 向通道发数据
	go func() {
		for {
			// 模拟生成大数据，每个数据块大小为 1 KB
			dataBlocks := []GameMetrics{
				{Game: "Example Game 1", Players: 100, Score: 1000},
				{Game: "Example Game 2", Players: 200, Score: 2000},
				{Game: "Example Game 3", Players: 300, Score: 3000},
				{Game: "Example Game 3", Players: 300, Score: 3000}, // 加了防抖机制， 两条数据相同，只发送一条
				// 添加更多数据块
			}

			for _, metrics := range dataBlocks {
				// 序列化为 JSON
				metricsJSON, err := json.Marshal(metrics)
				if err != nil {
					log.Println("Error marshaling JSON:", err)
					continue
				}

				// 发布消息到 Redis 频道
				err = redisClient.Client.Publish(ctx, "game_metrics_channel", string(metricsJSON)).Err()
				if err != nil {
					log.Println("Error publishing message:", err)
				}
				time.Sleep(time.Second * 1) // 每秒发布一个数据块
			}
		}
	}()

	return r
}

// GameMetrics 定义游戏指标数据的结构体
type GameMetrics struct {
	Game    string `json:"game"`
	Players int    `json:"players"`
	Score   int    `json:"score"`
}
