package main

import "fmt"

func DeferFunc1(i int) (t int) {

	fmt.Println("t = ", t)

	defer func() {
		t *= 10
	}()

	return 5
}

func main() {
	//fmt.Println(DeferFunc1(10))
	fmt.Println("result ", testPanic())
	fmt.Println(222)
}

func testPanic() int {
	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		if err := recover(); err != nil {
			fmt.Printf("print err %v \n ", err)
		}
	}()

	defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()
	panic("异常触发")

	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()

	fmt.Println("不能执行")
	return 2
}
