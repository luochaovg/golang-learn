package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
)

// gin helloword

// 定义接收数据的结构体
type Login struct {
	// binding:"required" 修饰的字段， 若为空值则报错，是必须字段
	User     string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间价 Logger(), Recovery()
	// 核心 r:=gin.Default 也可以创建不带中间价的路由
	r := gin.Default()

	// 路由组1，处理GET请求
	v1 := r.Group("/v1")
	// {}书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}

	// 路由组v2
	v2 := r.Group("/v2")
	v2.POST("/register", register)
	v2.POST("submit", submit)

	// 所有路由的规则构造一颗前缀树
	r.POST("/", login)
	r.POST("/search", login)
	r.POST("/support", login)
	r.POST("/blog/:post", login)
	r.POST("/contact", login)

	// 2.绑定路由规则，执行的函数
	// gin.Context 封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.POST("/post", postDemo)
	r.PUT("/put")
	r.DELETE("/delete")

	// 3.api 参数
	r.GET("/user/:name/*action", func(c *gin.Context) {

		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})

	// 4.URL参数
	// DefaultQuery()若参数不存在，返回默认值。 Query()若不存在,返回空串
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "zhagnsan")
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	// 5.表单参数
	// post , 四种格式
	// application/json
	// application/x-www-form-urlencoded
	// application/xml
	// multipart/form-data
	r.POST("/form", func(c *gin.Context) {
		// 表单参数设置默认值
		type1 := c.DefaultPostForm("type", "alert")

		// 接收其他的
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 多选
		hobbys := c.PostFormArray("hobdy")
		c.String(http.StatusOK,
			fmt.Sprintf("type is %s, useranme is %s, password is %s, hobby is %v", type1, username, password, hobbys))
	})

	// 5.1 单文件提交
	// 限制表单上传大小， 默认32MB
	r.MaxMultipartMemory = 8 << 20 // 限制文件上传大小8MB
	r.POST("/upload", func(c *gin.Context) {
		// 表单取文件
		file, _ := c.FormFile("avtor")
		log.Println(file.Filename)
		// 传到项目根目录， 名字就用本身
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s upload", file.Filename))
	})

	// 5.2上传多个文件
	r.POST("/uploads", func(c *gin.Context) {
		form, e := c.MultipartForm()
		if e != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", e.Error()))
		}

		// 获取所有图片
		files := form.File["avtor"]
		// 遍历所有图片
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", e.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})

	// gin数据解析绑定
	// JSON 绑定到结构体
	r.POST("/loginJson", func(c *gin.Context) {
		// 声明接收的变量
		var userJson Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&userJson); err != nil {
			// 返回错误信息
			// gin.H 封装了生成了json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 判断用户名密码是否正确
		if userJson.User != "root" || userJson.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 304,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": 200})
	})

	// 表单 绑定到结构体
	r.POST("/loginForm", func(c *gin.Context) {
		var userForm Login
		// bind()默认解析并绑定form格式
		// 根据请求头cotent-type自动推断
		if err := c.Bind(&userForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": userForm.User,
			"password": userForm.Password,
		})

	})

	// URL 绑定到结构体
	r.GET("/loginUrl/:username/:password", func(c *gin.Context) {
		var userUrl Login
		if err := c.ShouldBindUri(&userUrl); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": userUrl.User,
			"password": userUrl.Password,
		})
	})

	r.GET("/loginUrlParams", func(c *gin.Context) {
		var userUrl Login
		if err := c.ShouldBindQuery(&userUrl); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": userUrl.User,
			"password": userUrl.Password,
		})
	})

	// 多种响应方式
	// 1.json
	r.GET("/respJson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "respon json"})
		return
	})
	// 2.结构体
	r.GET("/respStruct", func(c *gin.Context) {
		var msg struct {
			User, Password string
		}
		msg.User = "admin"
		msg.Password = "root"

		c.JSON(http.StatusOK, msg)
	})
	// 3.xml
	r.GET("/respXml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"user":     "admin",
			"password": "root",
		})
	})
	// 4.yaml
	r.GET("/respYaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"user": "Admin"})
	})
	// 5.protobuf格式，谷歌开发的高效存储读取的工具
	r.GET("/respProtobuf", func(c *gin.Context) {
		reps := []int64{
			int64(1), int64(2),
		}
		label := "label"

		// protobuf 格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	// 设置listen 端口 , 默认：8080
	r.Run(":8000")
}

func postDemo(c *gin.Context) {

}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "Jack")
	c.String(http.StatusOK, name)
}
func register(c *gin.Context) {
	name := c.DefaultQuery("name", "zhangsahan")
	c.String(http.StatusOK, name)
}
func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lisi")
	c.String(http.StatusOK, name)
}
