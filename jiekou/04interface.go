package main

import (
	"fmt"
)

// 使用值接收者和指针接收者的区别？
// 使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
// 指针接收者实现接口只能存结构体指针类型变量

// 同一个结构体可以实现多个接口
// 多个类型可以实现同一个接口
// 一个类型可以实现多个接口

type animal interface {
	move()
	eat(food string)
}

type cat struct {
	name string
	feet int
}

type dog struct {
	name string
	feet int
}

// 使用值接收者实现了接口的所有方法
//func (c cat) move() {
//	fmt.Println("走猫步")
//}
//
//func (c cat) eat(food string) {
//	fmt.Printf("猫吃%s ... \n", food)
//}

func (d dog) move() {

}

func (d dog) eat(foot string) {
	fmt.Printf("dog name %s eat %s \n", d.name, foot)
}

// 使用指针接收者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("走猫步")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s ... \n", food)
}

func main() {

	var a1 animal

	var c1 cat = cat{
		name: "tom",
		feet: 4,
	}

	c2 := &cat{ // * cat
		name: "jialaolian",
		feet: 4,
	}

	a1 = &c1 // 实现animal 这个接口的是cat的指针类型
	fmt.Printf("%T %p\n", a1, a1)

	a1 = c2
	fmt.Printf("%T %p\n", a1, a1)

	//fmt.Println(a1)

	var d1 dog = dog{name: "xx", feet: 4}
	d1.eat("gutou")

	var a2 animal
	a2 = d1
	fmt.Printf("%T %p %#v \n", a2, a2, a2)

}
