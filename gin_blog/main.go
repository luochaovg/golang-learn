package main

import (
	"gin_blog/controller"
	"gin_blog/dao/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// parseTime=true  数据库时间字段解析为Go的时间字段，默认不会开
	dns := "root:123456@tcp(192.168.158.88:3306)/blogger?parseTime=true"

	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	// 加载静态文件
	//r.Static("/static/", "./static")
	//加载模版
	//r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "gin blog",
		})
	})

	r.GET("/index", controller.IndexHandler)
	r.GET("/category", controller.CategoryList)
	r.GET("/detail", controller.GetArticleDetail)

	_ = r.Run(":8000")
}
