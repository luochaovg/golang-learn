package db

import (
	"testing"
	"time"

	"gin_blog/model"
)

func init() {
	// parseTime=true  数据库时间字段解析为Go的时间字段，默认不会开
	dns := "root:lc910112@tcp(192.168.158.88:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	// 插入文章
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "this a test ak dkdkdkddkddkd"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Summary = `使用mysql的时间字段遇到如下两个问题
	1.使用go-sql-driver来连接mysql数据库，获取的时区默认是UTC +0的，与本地的东八区是有区别，在业务处理中会出现问题
	2.获取mysql中的日期，是string类型，需要在代码中用time.Parse进行转化`
	article.ArticleInfo.Title = "GOLANG 连接Mysql的时区问题"
	article.ArticleInfo.Username = "lhw"
	article.ArticleInfo.ViewCount = 1
	article.Category.CategoryId = 1

	articleId, err := InsertArticle(article)
	if err != nil {
		t.Errorf("insert article failed, err:%v\n", err)
		return
	}

	t.Logf("insert article succ, articleId:%d\n", articleId)
	//InsertArticle(article *model.ArticleDetail) (articleId int64, err error)
}

func TestGetArticleList(t *testing.T) {
	// 获得文章列表
	articleList, err := GetArticleList(0, 15)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Log(articleList)
	t.Logf("get article succ, len:%d\n", len(articleList))
}

func TestGetArticleInfo(t *testing.T) {
	// 获得文章详情
	articleInfo, err := GetArticleDetail(1)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}

	t.Logf("get article succ, article:%#v\n", articleInfo)
}

func TestGetRelativeArticle(t *testing.T) {
	// 获得相关文章
	articleList, err := GetRelativeArticle(7)
	if err != nil {
		t.Errorf("get relative article failed, err:%v\n", err)
		return
	}

	for _, v := range articleList {
		t.Logf("id:%d title:%s\n", v.ArticleId, v.Title)
	}
}

func TestGetPrevArticleById(t *testing.T) {
	// 测试上一篇
	articelInfo, err := GetPrevArticleById(6)
	if err != nil {
		t.Errorf("get prev article failed, err:%v\n", err)
		return
	}

	t.Logf("artice info:%#v", articelInfo)
}

func TestGetNextArticleById(t *testing.T) {
	// 获取下一篇
	articelInfo, err := GetNextArticleById(6)
	if err != nil {
		t.Errorf("get prev article failed, err:%v\n", err)
		return
	}

	t.Logf("artice info:%#v", articelInfo)
}
