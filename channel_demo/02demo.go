package main

import (
	"fmt"
	"sync"
	"time"
)

// 通道引用类型

// channel 练习
// 1.启动一个goroutine，生产100个数发送到ch1
// 2.启动一个goroutine，从ch1中取值，计算其平方放到ch2中
// 3.在main中，从ch2中取值打印

// 注意： 往通道存值 -> 发送  ， 往通道取值 -> 接收

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		ch1 <- i
	}

	close(ch1) // 关闭通道不能在写了， 但是可以读
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	//for x := range ch1 {
	//	ch2 <- x * x
	//}

	for {
		x, ok := <-ch1
		if !ok {
			break
		}

		ch2 <- x * x
	}

	once.Do(func() { close(ch2) }) // 确保某个操作只执行一次
}

func main() {

	a := make(chan int, 100)
	b := make(chan int, 100)

	wg.Add(3)

	go f1(a)    // 生产
	go f2(a, b) // 消费1
	go f2(a, b) // 消费2

	wg.Wait()

	for ret := range b {
		fmt.Println(ret)
	}

	time.Tick(time.Second)
}
