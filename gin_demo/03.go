package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件
// gin可以构建中间件，但它只对注册过的路由函数起作用
// 对于分组路由，欠她使用中间件，可以限定中间件的作用范围
// 中间件分为全局中间件，单个路由中间件和群组中间件
// TODO gin 中间件必须是一个gin.HandlerFunc类型

// 定义 全局中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("全局中间件开始执行")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("isLogin", false)
		// 执行函数
		c.Next()

		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("全局中间件执行完毕", status)

		end := time.Since(start)
		fmt.Println("time:", end)
	}
}

// 另外一种定义中间件的方法
func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	//统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)
}

// 局部中间件
func partMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("局部中间件开始执行")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("token", "111111")
		// 执行函数
		c.Next()

		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("局部件执行完毕", status)

		end := time.Since(start)
		fmt.Println("time:", end)
	}
}

func main() {
	r := gin.Default()

	// 注册中间件
	r.Use(MiddleWare())
	{
		r.GET("/middle_test", func(c *gin.Context) {
			// 取值
			isLogin, _ := c.Get("isLogin")
			fmt.Println("is Login", isLogin)

			c.JSON(http.StatusOK, gin.H{"is_login": isLogin})
		})

		// 跟路由后面是定义的局部中间件
		r.GET("/middle_test2", partMiddleWare(), func(c *gin.Context) {
			// 取值
			token, _ := c.Get("token")
			fmt.Println("Token", token)

			c.JSON(http.StatusOK, gin.H{"token": token})
		})
	}

	// 路由组，中间件
	g1 := r.Group("/g1", partMiddleWare())
	{
		g1.GET("/middle1", func(c *gin.Context) {
			// 取值
			token, _ := c.Get("token")
			fmt.Println("Token", token)

			c.JSON(http.StatusOK, gin.H{"token": token})
		})
		g1.GET("/middle2", func(c *gin.Context) {
			// 取值
			token, _ := c.Get("token")
			fmt.Println("Token", token)

			c.JSON(http.StatusOK, gin.H{"token": token})
		})
	}

	s1 := r.Group("/shoping", myTime) // 作为参数，另外一种形式
	{
		s1.GET("/index", shopIndexHandler)
		s1.GET("/home", homeIndexHandler)
	}
	r.Run(":8000")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{"msg": "shop index"})
}
func homeIndexHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
	c.JSON(http.StatusOK, gin.H{"msg": "home index"})
}
