package main

import (
	"runtime"
	"sync"
)
import "fmt"

var wg sync.WaitGroup

// 利用有缓冲channel+sync
func main() {
	ch := make(chan bool, 3)

	// 往ch里面发送任务
	w := 10000
	for i := 0; i <= w; i++ {
		wg.Add(1)
		ch <- true

		go work(ch, i)
	}

	wg.Wait()
}

func work(ch chan bool, i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())

	<-ch
	wg.Done()

}
