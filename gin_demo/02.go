package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// HTML 渲染 / 重定向 / 同步异步
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "02 Hello world")
	})

	// 加载模版文件
	r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLGlob("templates/index.tmpl")

	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 最终是json将title替换
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "我的标题",
		})
	})

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		// 支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		//c.Redirect(http.StatusMovedPermanently, "/index")
	})

	// 异步
	// TODO goroutine 机制可以方便实现异步处理
	// TODO 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
	r.GET("/loginSync", func(c *gin.Context) {
		// 需要另外一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(time.Second * 3)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	// 同步
	r.GET("login_tongbu", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

	r.Run(":8000")
}
