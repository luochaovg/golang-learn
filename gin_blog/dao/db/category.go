package db

import (
	"gin_blog/model"
	"github.com/jmoiron/sqlx"
)

// 插入分类
func InsertCategory(category *model.Category) (cateGoryId int64, err error) {
	sqlStr := "insert into category(category_name, category_no)values(?,?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}

	cateGoryId, err = result.LastInsertId()
	return
}

// 根据ids 获取多分类
func GetCategoryListByIds(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlStr, args, err := sqlx.In("select id, category_name, category_no from category where id in(?)", categoryIds)
	if err != nil {
		return
	}

	// 传入时必须将切片展开
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

// 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlStr)
	return
}

// 根据单个id 获取分类
func GetGategoryById(id int64) (category *model.Category, err error) {
	sqlStr := "select id, category_name, category_no from category where id=?"

	err = DB.Get(&category, sqlStr, id)

	return
}
