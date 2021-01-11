package main

import "fmt"

type animal interface {
	move()
	eat(food string)
}

type cat struct {
	name string
	feet int
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s ... \n", food)
}

func main() {

	var a1 animal // 定义一个接口类型变量

	bc := cat{
		name: "蓝猫",
		feet: 4,
	}

	// 接口保存分为两部分 （动态类型/动态值） ， 这样就实现了接口变量能够存储不同的值
	// 存类型： main.cat  存值：name:"蓝猫"，feet:4
	a1 = bc
	a1.move()
	a1.eat("老鼠")
	fmt.Println(a1)
	fmt.Printf("%T \n", a1)

}
