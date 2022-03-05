package main

import (
	"fmt"
	"sort"
)

type SortSliceTest struct {
	Num  int
	Name string
}

func main() {
	arrs := InitArrs()
	sort.Slice(arrs, func(i, j int) bool {
		return arrs[i].Num > arrs[j].Num
	})

	for k, v := range arrs {
		fmt.Println("index", k, "value", v)
	}
}

func InitArrs() (arrs []SortSliceTest) {
	arr1 := SortSliceTest{
		Num:  3,
		Name: "arr1",
	}

	arr2 := SortSliceTest{
		Num:  1,
		Name: "arr2",
	}

	arr3 := SortSliceTest{
		Num:  5,
		Name: "arr3",
	}

	arr4 := SortSliceTest{
		Num:  2,
		Name: "arr4",
	}

	arrs = append(arrs, arr1, arr2, arr3, arr4)
	return
}


