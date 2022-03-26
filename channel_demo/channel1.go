package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func busi(ch chan bool, i int) {

	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
}

func main() {
	//模拟用户需求业务的数量
	ch := make(chan bool, 3)
	task_cnt := math.MaxInt64
	//task_cnt := 10

	for i := 0; i < task_cnt; i++ {

		ch <- true

		go busi(ch, i)
	}

}

func fas() {
	var ch = make(chan int)
	for i := 1; i < 75; i++ {
		go print2(i)
		ch <- 0
	}
	time.Sleep(5 * time.Second)
}

// 存在以下代码，请问如何实现print函数可以顺序输出1~75，要求不使用锁，只能使用channel
func print2(i int) {
	<-ch
	fmt.Printf("i is %d \n", i)
}

func printn(n int) {
	ch := make(chan int, 1)
	for i := 0; i <= n; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Printf("i is %d \n", v)
		default:
		}
	}
}
