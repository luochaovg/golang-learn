package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

// 并发安全
// 说在前面的话：这是一个进阶知识点。
//
//在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，
// 例如只加载一次配置文件、只关闭一次通道等。
//
//Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。
//
//sync.Once只有一个Do方法，其签名如下：
//
//func (o *Once) Do(f func()) {}
//

var wg sync.WaitGroup
var once sync.Once // 并发安全

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}

	f := func() {
		close(ch2)
	}
	once.Do(f)
}

func main() {
	// go run 04demo.go 2> trace.out
	//  go tool trace trace.out
	trace.Start(os.Stderr)
	defer trace.Stop()

	a := make(chan int, 100)
	b := make(chan int, 100)

	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()

	for ret := range b {
		fmt.Println(ret)
	}
}
