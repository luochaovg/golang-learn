package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context?
var wg2 sync.WaitGroup
var ch chan bool = make(chan bool, 1)

func f1() {
	defer wg2.Done()

LOOP:
	for {
		fmt.Println("hello")
		time.Sleep(time.Second)
		select {
		case <-ch:
			break LOOP
		default:
		}
	}
}

func main() {
	wg2.Add(1)
	go f1()

	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出?
	ch <- true

	wg2.Wait()
}
