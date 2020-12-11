package main

import (
	"fmt"
	"sync"
	"time"
)

// https://www.liwenzhou.com/posts/Go/go_context/

// 为什么需要context?

var wg sync.WaitGroup
var notify bool

func f() {
	defer wg.Done()

	for {
		fmt.Println("hello")
		time.Sleep(time.Second)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()

	time.Sleep(time.Second * 10)
	// 如何通知子goroutine退出?
	notify = true

	wg.Wait()
}
