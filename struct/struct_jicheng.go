package main

import "fmt"

// 结构体模拟实现面向对象继承

type animal struct {
	name string
}

// 给animail 实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s 会动 \n", a.name)
}

// 狗子
type dog struct {
	feet   uint8
	animal // animal 拥有的方法，dog此时也有了， 变向实现了继承
}

// 给dog实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s wang wang \n", d.name)
}

func main() {

	d1 := dog{
		feet: 4,
		animal: animal{
			name: "狗子",
		},
	}

	fmt.Println(d1)

	d1.move()
	d1.wang()

}
