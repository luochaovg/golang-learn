package db

import (
	"database/sql"
	"fmt"

	"gin_blog/model"
	_ "github.com/go-sql-driver/mysql"
)

// 插入文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}

	sqlstr := `insert into 
					article(content, summary, title, username, 
						category_id, view_count, comment_count)
				values(?, ?, ?, ?, ?, ?, ?)`

	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	// 返回插入的id
	articleId, err = result.LastInsertId()
	return
}

// 获取文章列表
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
					order by create_time desc
					limit ?, ?`
	// limit 从哪条开始，步长
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}

// 根据分类id,查询这一类文章
func GetArticleListByCategoryId(categoryId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {

	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
						and
						category_id = ?
					order by create_time desc
					limit ?, ?`

	err = DB.Select(&articleList, sqlstr, categoryId, pageNum, pageSize)
	return
}

// 根据id获取文章详情
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", articleId)
		return
	}

	articleDetail = &model.ArticleDetail{}
	sqlstr := `select 
							id, summary, title, view_count, content,
							 create_time, comment_count, username, category_id
						from 
							article 
						where 
							id = ?
						and
							status = 1
						`

	err = DB.Get(articleDetail, sqlstr, articleId)
	fmt.Printf("article_info:%#v\n", articleDetail)
	return
}

// 相关文章
func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	// 获取该文章的分类类别，展示其中的10篇
	var categoryId int64
	sqlstr := "select category_id from article where id=?"
	err = DB.Get(&categoryId, sqlstr, articleId)
	if err != nil {
		return
	}

	sqlstr = "select id, title from article where category_id=? and id !=?  limit 10"
	err = DB.Select(&articleList, sqlstr, categoryId, articleId)
	return
}

// 获取指定id的上一篇
func GetPrevArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlstr := "select id, title from article where id < ? order by id desc limit 1"
	err = DB.Get(info, sqlstr, articleId)
	if err != nil {
		return
	}

	return
}

// 获取指定id的下一篇
func GetNextArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		// 没有获取到下一篇，将文章ID置为-1
		ArticleId: -1,
	}
	sqlstr := "select id, title from article where id > ? order by id asc limit 1"
	err = DB.Get(info, sqlstr, articleId)
	if err != nil {
		return
	}

	return
}

func IsArticleExist(articleId int64) (exists bool, err error) {
	// 判断文章id是否存在
	var article_id int64
	sqlstr := "select id from article where id=?"
	err = DB.Get(&article_id, sqlstr, articleId)
	//fmt.Println("----",article_id)
	if err == sql.ErrNoRows {
		exists = false
		return
	}

	if err != nil {
		return
	}

	exists = true
	return
}
