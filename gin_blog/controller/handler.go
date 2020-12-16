package controller

import (
	"fmt"
	"gin_blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 访问主页的控制器
func IndexHandler(c *gin.Context) {
	// 从service 取数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	categoryList, err := service.GetAllCategoryByList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// 定义返回变量
	//var data map[string]interface{} = make(map[string]interface{}, 10)
	//data["article_list"] = articleRecordList
	//data["category_list"] = categoryList

	// gin.H 本质上是一个map
	c.JSON(http.StatusOK, gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
	return
}

// 首页文章分类信息列表
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category_id": categoryId,
	})
}

// 获取文章详情
func GetArticleDetail(c *gin.Context) {
	articleIdStr := c.Query("id")
	id, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err,
		})
		return
	}

	articleDetail, err := service.GetArticleDetailById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article_detail": articleDetail,
	})
}
