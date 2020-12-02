package base

import (
	"fmt"
)

func main() {
	//fmt.Println()
	//fmt.Print()
	//fmt.Printf()

	//var m1 = make(map[string]int, 2)
	//m1["age"] = 100
	//m1["sex"] = 1
	//
	//fmt.Printf("%v \n ", m1)
	//fmt.Printf("%#v \n", m1)
	//fmt.Println(len(m1))

	//var s string
	//fmt.Scan(&s) // 修改s值
	//fmt.Println("输入内容：", s)

	//var (
	//	name  string
	//	age   int
	//	class string
	//)

	name := "luochao"
	t := 'l'
	fmt.Printf("%T \n", t)
	for _, v := range name {
		if v == t {
			fmt.Printf("%c \n", v)
		}

	}
}
