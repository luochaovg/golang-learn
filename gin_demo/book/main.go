package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	// 加载页面
	r.LoadHTMLGlob("./templates/*")

	r.GET("/book/list", bookListHandler)

	r.Run(":8000")
}

func bookListHandler(c *gin.Context) {
	bookList, err := QueryAllBook()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}

	c.HTML(http.StatusOK, "list.html", gin.H{
		"code": 0,
		"msg":  "Success",
		"data": bookList,
	})
}
