package service

import (
	"gin_blog/dao/db"
	"gin_blog/model"
)

// 获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1，获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}

	if len(articleInfoList) <= 0 {
		return
	}

	// 2，获取文章对应的分类
	categoryIds := GetCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryListByIds(categoryIds)
	if err != nil {
		return
	}

	// 3,组合数据
	for _, article := range articleInfoList {
		// 根据当前的文章，生产结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}

		// 取出文章分类id
		categoryId := article.CategoryId
		// 文章取出分类
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}

		articleRecordList = append(articleRecordList, articleRecord)
	}

	return
}

// 根据多个文章的id,获取多个分类id的集合
func GetCategoryIds(articleList []*model.ArticleInfo) (categoryIds []int64) {
LABEL:
	for _, article := range articleList {
		// 从当前文章取出分类id
		categoryId := article.CategoryId

		// 去重
		for _, id := range categoryIds {
			// 分类列表去重
			if id == categoryId {
				continue LABEL
			}
		}

		categoryIds = append(categoryIds, categoryId)
	}
	return
}

// 获取文章详情
func GetArticleDetailById(id int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(id)
	if err != nil {
		return
	}

	return
}
