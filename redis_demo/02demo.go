package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"time"
)

var (
	rdb2 *redis.Client
)
var ctx2 = context.TODO()

func init() {
	rdb2 = redis.NewClient(&redis.Options{
		Addr:     "172.16.13.143:6379",
		Password: "lc910112",
		DB:       0,
		PoolSize: 100,
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := rdb2.Ping(timeout).Result()
	if err != nil {
		panic(err)
	}
}


// 保存string 类型的数据入redis
func SaveStringToRedis(key, data string, expireTime time.Duration) (b bool, err error) {
	err = rdb2.Set(ctx2 ,key, data, expireTime).Err()
	if err != nil {
		zap.L().Error("save string to redis failer",
			zap.String("key:", key), zap.String("data", data), zap.Error(err))
		return false, err
	}

	return true, nil
}

// 从redis获取string类型的数据
func GetStringByRedis(key string) (value string, err error) {
	value, err = rdb2.Get(ctx2, key).Result()

	if err != nil {
		return value, err
	}
	return value, nil
}

// 设置redis key 过期时间
func SetExpire(key string, expireTime time.Duration) (err error) {
	err = rdb2.Expire(ctx2, key, expireTime).Err()
	if err != nil {
		zap.L().Error("set redis key expire failed", zap.String("key", key), zap.Error(err))
		return
	}
	return
}

// 判断redis key值是否存在
func IsKeyExist(key string) bool {
	if rdb2.Exists(ctx2, key).Val() < 1 {
		return false
	}
	return true
}

type CacheMetricesByDateList struct {
	GameCode string
	Filter []string
	CacheUsed bool
	Metircs string
	Expire time.Duration
	Date string
}

func NewCacheMetricesByDateList(gameCode , metircs string, filter []string , isCache bool) *CacheMetricesByDateList  {
	return  &CacheMetricesByDateList{
		GameCode:  gameCode,
		Filter:    filter,
		CacheUsed: isCache,
		Metircs: metircs,
		Date: "2022",
	}
}

// 业务逻辑处理
func (a *CacheMetricesByDateList)RevenueMetricLogic()(interface{}, error)  {

	fmt.Println("aaa", a.GameCode)
	//fmt.Println("ga", gameCode)

	return  a.GameCode , nil
}

// 业务逻辑函数
type MetricesLogicFunc func() (interface{}, error)

// 缓存
func (a *CacheMetricesByDateList)FunWapper(f MetricesLogicFunc)(interface{}, error)  {
	cacheKey := a.GameCode+a.Metircs

	// 查缓存
	if IsKeyExist(cacheKey) {
		val , _ := GetStringByRedis(cacheKey)
		if len(val) > 0 {
			return  val, nil
		}
	}

	// 业务逻辑
	a.Date = "2010"
	rst , err := f()
	if err != nil {
		fmt.Println(err)
		return nil , err
	}

	// 缓存
	_, err = SaveStringToRedis(cacheKey, rst.(string), a.Expire)
	if err != nil {
		fmt.Println(err)
		return rst, nil
	}

	return  rst, nil
}

func main() {

	//b, err := SaveStringToRedis("tst", "luochao", time.Hour)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(b)

	api := &CacheMetricesByDateList{
		GameCode:  "test",
		Filter:    []string{"aa"},
		CacheUsed: true,
		Metircs:   "revunue",
		Expire:    60*time.Second,
		Date: "2025",
	}

	val, err:= api.FunWapper(api.RevenueMetricLogic)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
