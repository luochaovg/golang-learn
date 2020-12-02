package base

import (
	"fmt"
	"strings"
)

// 函数闭包
func f1(f func()) {
	fmt.Println("f1")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
		fmt.Println("hello")
	}

	return tmp
}

func addr(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包：一个函数除了可以内部变量外还可以引用外部作用域的变量

// 1, 函数可以作为返回值
// 2，函数内部查找变量的顺序， 先在自己内部找，找不到往外层找

// 闭包 demo1
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

// 闭包 demo 2
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sum := func(i int) int {
		base -= i
		return base
	}

	return add, sum
}

func main() {

	//f1 := func() {
	//	fmt.Println("aaa")
	//}
	//f1()
	//
	//// 函数后面加括号，立即执行
	//func(x, y int) {
	//	fmt.Println(x + y)
	//}(100, 200)
	//
	//ret := addr(100)
	//
	//ret2 := ret(200)
	//fmt.Println(ret2)

	//
	//ret := f3(f2, 100, 299)
	//f1(ret)
	//fmt.Printf("%T \n", ret)

	//jpgFunc := makeSuffixFunc(".jpg")
	//fmt.Printf("%T \n ", jpgFunc)
	//name := jpgFunc("test.jpg")
	//fmt.Println(name)

	add, sub := calc(10)
	//fmt.Println(add(1), sub(2))
	fmt.Println(add(3), sub(4))
}
