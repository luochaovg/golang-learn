package db

import "testing"

func init() {
	dns := "root:123456@tcp(192.168.158.88:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}
	for _, v := range categoryList {
		t.Logf("id:%d, category:%#v\n", v.CategoryId, v)
	}
}

func TestGetCategoryListByIds(t *testing.T) {
	categoryIds := []int64{1, 2, 3}
	//var categoryIds []int64
	//categoryIds = append(categoryIds, 1, 2, 3)

	categoryList, err := GetCategoryListByIds(categoryIds)
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}

	// 获取的长度和传入的id长度一致
	if len(categoryList) != len(categoryIds) {
		t.Errorf("get category list failed, len of categorylist:%d ids len:%d\n",
			len(categoryList), len(categoryIds))
	}

	for _, v := range categoryList {
		t.Logf("id:%d, category:%#v\n", v.CategoryId, v)
	}
}
