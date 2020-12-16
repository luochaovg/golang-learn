package main

import "fmt"

// 接口是一种类型, 是一种特殊的类型，它规定了变量有哪些的方法
// 在编程中会遇到以下场景
// 我不关心一个变量是什么类型， 我只关心能调用它的方法

// 引出接口实例
type cat struct{}
type dog struct{}
type person struct{}

// 定义一个能叫的类型
type speaker interface {
	//TODO 只要实现了speak 方法的变量都是speaker类型
	speak()
}

/**
接口的定义
type 接口名 interface {
	方法名1(参数1，参数2...) (返回值1， 返回值2..)
	方法名2(参数1，参数2...) (返回值1， 返回值2..)
	...
}
用来给变量\参数\返回值设置类型

接口的实现
一个变量如果实现了接口中规定的所有的方法，那么这个变量就实现了这个接口，可以称为这个接口类型的变量
*/

func (c cat) speak() {
	fmt.Println("mimimi")
}
func (d dog) speak() {
	fmt.Println("wangwangwang")
}
func (p person) speak() {
	fmt.Println("aaaaa")
}

func da(x speaker) {
	// 接收一个参数，传进来什么，我就大什么
	x.speak() // 挨打就叫
}

func main() {

	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker // 定义一个接口类型speaker, 的变量ss
	ss = c1

	fmt.Println(ss)

}
