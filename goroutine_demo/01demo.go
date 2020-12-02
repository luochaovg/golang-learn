package main

import (
	"fmt"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// goroutine 什么时候结束
//  goroutine 对应的函数结束了， goroutine结束了
//  main 函数执行完了，有main函数创建的那些goroutine都结束了

// 程序启动之后会创建一个main goroutine 去执行
func main() {
	// go hello(1) // 开启一个单独的goroutine去执行hello 函数（任务）

	for i := 0; i < 10; i++ {
		//go hello(i)
		go func(i int) {
			fmt.Println(i) // 用的是函数参数的那个i， 不是外面的i
		}(i)
	}

	fmt.Println("mail")
	// time.Sleep(time.Second)
	// main 函数结束了，由main函数启动的goroutine也都结束了
}
