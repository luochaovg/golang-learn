package service

import (
	"gin_blog/dao/db"
	"gin_blog/model"
)

// 获取所有分类
func GetAllCategoryByList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}
