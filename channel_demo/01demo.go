package main

import (
	"fmt"
	"sync"
)

// channel
// goroutine 是go程序并发的执行体，channel 就是他们的链接
// channel 是一种特定的类型

var slice []string
var ch chan int // 需要指定通道中元素的类型
var n int
var wg sync.WaitGroup

// make : map slice channel

func noBufChannel() {
	fmt.Printf("%T, %#v \n", slice, slice)

	fmt.Println(ch) // nil

	// 通道必须使用make函数初始化才能使用, 不带缓冲区
	ch = make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch
		fmt.Println("后台goroutine从通道中渠道了", x)
	}()

	ch <- 10 // 如果没有接收goroutine将会卡住了报错
	fmt.Println("发送到通道中", ch)

	wg.Wait()
}

func bufChannel() {
	// 带缓冲区的通道
	ch1 := make(chan int, 10)
	ch1 <- 10
	fmt.Println("10发送到通道中", ch1)

	ch1 <- 20
	fmt.Println("20发送到通道中", ch1)

	x := <-ch1
	fmt.Println(x)

	close(ch1)
}
func main() {

	bufChannel()
	// 发送
	//ch1 <- 10 // 把10发送到通道
	//fmt.Printf("%T, %#v \n", ch, ch)

	// 接收
	//n = <-ch
	//<-ch // 丢弃

	// 关闭
	//close(ch)
}
