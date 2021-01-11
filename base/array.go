package main

import "fmt"

func main() {

	// 数组 存放元素的容器， 必须指定存放元素的类型和容量（长度）
	// 数组的长度是数组类型的一部分
	// 数组是值类型
	var a1 [2]int
	var a2 [3]bool

	fmt.Println(a1, a2, a2[1])

	// 数组的初始化1
	a1 = [2]int{1, 3}

	a1[1] = 8
	fmt.Println(a1)

	// 数组的初始 2
	a3 := [3]int{5, 6, 2}

	// 根据初始值自动推断数组长度
	a4 := [...]int{5, 6, 2, 5}
	fmt.Println(a3, a4)

	c1 := [...]int{3, 5, 6, 7}
	c2 := c1[0:2]
	fmt.Println(c1, c2)

	// 数组的初始 3, 根据索引初始化，没有指定索引默认0
	a5 := [4]string{0: "你", 3: "fa"}
	fmt.Println(a5, len(a5))

	for i, v := range a5 {

		if v == "" {
			break
		}
		fmt.Println(i, v)
	}

	// 多纬数组
	a6 := [3][2]int{
		{2, 3},
		{1, 3},
		{2, 6},
	}

	fmt.Println(a6)

	for _, v := range a6 {
		for _, c := range v {
			fmt.Println(c)
		}
	}

	b1 := [3]int{6, 7, 8}
	b2 := b1
	b1[0] = 9
	fmt.Println(b1, b2, a1)

	// 数组支持"=="， "!=" 操作符， 因为内存总是被初始化过
	// [n]*T 表示指针数组， *[n]T表示数组指针

}
