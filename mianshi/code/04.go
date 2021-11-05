package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type MyInt1 int   // 创建新类型
type MyInt2 = int // 创建类型别名

func main() {
	p := new(Person)
	p.Name = "Luochao"
	(*p).Age = 23

	fmt.Println(p)

	var i int = 0
	var i1 MyInt1 = 0
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
