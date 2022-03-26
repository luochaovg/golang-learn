package main

import "sync"
import "fmt"

var wg sync.WaitGroup

// 利用无缓冲channel与任务发送/执行分离方式
func main() {
	ch := make(chan int)

	// 开三个goroutine干活
	cnt := 3
	for i := 0; i < cnt; i++ {
		go work(ch)
	}

	// 往ch里面发送任务
	w := 1000
	for i := 0; i <= w; i++ {
		wg.Add(1)
		ch <- i
	}

	wg.Wait()
}

func work(ch chan int) {
	// 从ch读数据
	for v := range ch {
		fmt.Println("int:", v)
		wg.Done()
	}
}
