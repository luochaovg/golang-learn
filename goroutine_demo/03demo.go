package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {

	// 默认跑满cpu核心数， 这里设定一核CPU的话， goroutine 排队执行
	runtime.GOMAXPROCS(1)

	fmt.Println(runtime.NumCPU())

	wg.Add(2)

	go a()
	go b()
	wg.Wait()
}

//Go语言中的操作系统线程和goroutine的关系：
//
//一个操作系统线程对应用户态多个goroutine。
//go程序可以同时使用多个操作系统线程。
//goroutine和OS线程是多对多的关系，即m:n。把m个goroutine分配给n个操作系统线程去执行
