package main

import "fmt"

/**
结构体遇到的问题
*/

type myInt int32

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
}

// 结构体的匿名字段， 字段少比较简单的场景， 不常用
type personNm struct {
	string
	int
}

// 结构体嵌套
type company struct {
	name string
	addr address
}

type address struct {
	province string
	city     string
}

func main() {

	var x int32
	x = 89
	fmt.Println(x)

	var y int32 = 34
	fmt.Println(y)

	z := int32(56)
	fmt.Println(z)

	a := int8(100) // (-128 - 127)
	fmt.Println(a)

	b := myInt(9)
	fmt.Println(b)

	//type person struct {
	//	name string
	//	age  int
	//}

	// 方法一
	var p person
	p.name = "law"
	p.age = 22

	// 方法二
	var p1 = person{
		name: "a",
		age:  0,
	}
	fmt.Println(p1)

	fmt.Printf("%T", p1)

	s1 := []int{1, 2, 3}
	a1 := [2]int{1, 3}
	m1 := map[string]int{"m1": 100, "m2": 200}

	fmt.Printf("%T,%T,%T \n", s1, a1, m1)

	pnm := personNm{
		string: "name",
		int:    34,
	}
	fmt.Println(pnm.string, pnm.int)

	p8 := person{
		name: "law",
		age:  30,
		address: address{
			province: "hunan",
			city:     "luodi",
		},
	}
	fmt.Println(p8, p8.age, p8.address.city, p8.province)

	var abc = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(abc)
}
