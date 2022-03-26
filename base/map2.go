package main

import "fmt"

// https://github.com/luochaovg/golang/blob/main/3%E3%80%81Map.md

type student struct {
	Name string
	Age  int
}

func main() {
	//定义map
	m := make(map[string]*student)

	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//将数组依次添加到map中
	// TODO foreach中，stu是结构体的一个拷贝副本，所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝。
	//for _, stu := range stus {
	//	m[stu.Name] = &stu
	//}

	for i := 0; i <= len(stus)-1; i++ {
		m[stus[i].Name] = &stus[i]
	}

	//打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
