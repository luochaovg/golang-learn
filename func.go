package main

import "fmt"

var m int = 100

// 命名返回值就相当于在函数中声明一个变量 ret
func sun(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func sum1(x int, y int) {
	fmt.Println(x + y)
}

func sum2() {
	fmt.Println("f2")
}

func sum3() int {
	return 3
}

// 使用命名返回值可以return后省略
func sum4(x int, y int) (ret int) {
	ret = x + y
	return
}

func sum5(x int, y int) (int, int) {

	return 1, 3
}

func sum55(x, y int) (a, b int) {
	a = x + y
	b = x - y
	return
}

// 参数类型简写, 当参数中连续两个参数的类型一致时， 我们可以将前面那个参数类型省略
func sum6(x, y int, m, n string, i, j bool) int {
	return x + y
}

// 可变传参 ， 可变长参数必须放在参数最后
func sum7(x string, y ...int) {
	fmt.Println(x, y) // y类型是切片
}

// go 语言函数中没有默认参数这个概念
// go语言中的函数传递的都是值 ctrl+c , ctrl+v

// 函数类型
func t1() {
	fmt.Println("hell law")
}

func t2() int {
	return 4
}

// 函数也可作为参数的类型
func t3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 函数可以作为返回值
func t4(x func() int) func(int, int) int {

	ret := func(a, b int) int {
		return a + b
	}

	return ret
}

func main() {
	//a, b := 6, 7
	//fmt.Println(sun(a, b))
	//fmt.Println(sum4(a, b))
	//
	//m, n := sum5(a, b)
	//fmt.Println(m, n)
	//
	//sum7("aa", 1, 4, 5)
	//
	//f := t1
	//c := t2
	//fmt.Printf("%T \n", f)
	//fmt.Printf("%T \n", c)
	//
	//t3(c)

	//defunc()
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())

	//a := 1
	//b := 2
	//defer calc("1", a, calc("10", a, b))
	//a = 0
	//defer calc("2", a, calc("20", a, b))
	//b = 1

	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Println(1)
}
func funcB() {
	defer func() { // defer 一定要在可能引发panic的语句之前定义
		err := recover() // recover 必须搭配defer 使用
		fmt.Println(err)
		fmt.Println("释放数据库链接")
	}()

	panic("数据库链接错误")
	fmt.Println(2)
}
func funcC() {
	fmt.Println(3)
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func defunc() {
	fmt.Println("start")

	defer fmt.Println("cccccc") // defer 把他后面的语句延迟函数即将返回的时候在执行
	defer fmt.Println("bbbb")   // 多个defer 后进先出

	fmt.Println(m)
	// go 语言中函数return不是原子操作， 在底层是分为两步执行
	// 1, 返回值赋值
	// defer
	// 2，真正的ret返回
	// 函数中如果存在defer, 那么defer 执行的时机实在1，2之间

	fmt.Println("end")
}

// 函数中查找变量的顺序
//	    1,先在函数内部查找
// 		2,如果找不到，往函数外面查找，一直找到全局变量

// 语句块作用域

// defer 只能执行一层函数

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5
}
func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数内x的值
	}(x)
	return 5 // 5
}

func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 6
}
