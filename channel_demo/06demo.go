package main

import (
	"fmt"
	"sync"
	"time"
)

//
//一个简易的work pool
//在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，
//控制goroutine的数量，防止goroutine泄漏和暴涨。
var wg sync.WaitGroup
var notifyChan = make(chan struct{}, 5)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		//fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second) // 执行业务逻辑
		fmt.Printf("worker:%d end job:%d\n", id, j)
		result <- j * 2

		notifyChan <- struct{}{} // 空结构体不占空间
	}

	//wg.Done()
}

//type cat struct{} // 声明类型
//var c1 = cat{}    // 实例化

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	go func() {
		// 5个任务 (producers) 往通道写数据
		for j := 1; j <= 5; j++ {
			jobs <- j
		}

		close(jobs)
	}()

	// 输出结果
	//for a := 1; a <= 5; a++ {
	//	r := <-result
	//	fmt.Println(r)
	//}

	//
	//for {
	//	// 如果result通道没有值，堵塞等待。 当result被关闭的时候 ok=false
	//	x, ok := <-result
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(x)
	//}

	// 开启3个goroutine从通道拿数据 (consumers 执行业务逻辑)
	for w := 0; w < 3; w++ {
		go worker(w, jobs, result)
	}

	//
	go func() {
		for i := 0; i < 5; i++ {
			<-notifyChan
		}

		close(result)
	}()

	// 输出结果     //无限循环从通道中读取数据
	for x := range result { // result 没有关闭会造成死锁
		fmt.Println(x)
	}

}
