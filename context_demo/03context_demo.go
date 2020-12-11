package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context?
// https://www.liwenzhou.com/posts/Go/go_context/

var wg3 sync.WaitGroup

func f2(ctx context.Context) {
	defer wg3.Done()

	go f3(ctx)

LOOP:
	for {
		fmt.Println("hello f2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func f3(ctx context.Context) {
	defer wg3.Done()

LOOP:
	for {
		fmt.Println("hello f3")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func main() {
	// context.Background() 根结点
	ctx, cancel := context.WithCancel(context.Background())

	wg3.Add(1)
	go f2(ctx)

	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出?
	cancel() // 通知 子goroutine退出

	wg3.Wait()
}
