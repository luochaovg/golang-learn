package main

import "fmt"

// 多个类型可以实现同一个接口

type car interface {
	run()
}

type fll struct {
	brand string
}

func (f fll) run() {
	fmt.Printf("%s 速度70 \n", f.brand)
}

type bsj struct {
	brand string
}

func (b bsj) run() {
	fmt.Printf("%s 速度100 \n", b.brand)
}

// drive 函数接收一个car 类型的变量
func drive(c car) {
	c.run()
}

func main() {

	var f1 = fll{brand: "法拉利"}
	var b1 = bsj{brand: "保时捷"}

	drive(f1)
	drive(b1)

}
