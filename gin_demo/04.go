package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("isLogin"); err == nil {
			if cookie == "yes" {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("isLogin", "yes", 60, "/",
			"localhost", false, true)
		c.String(http.StatusOK, "Login Success")
	})

	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})

	r.Run(":8000")
}
