package main

import (
	"fmt"
	"sync"
)

// 锁
// 互斥锁, 是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源
// 防止多个goroutine 同一时刻操作同一个资源

// sync.Mutex 是一个结构体， 是值类型，传函数传参数的时候要传指针

// sync.WaitGroup 也是一个结构体， 值类型，goroutine执行完在继续

var x = 0

var wg sync.WaitGroup

var lock sync.Mutex // 互斥锁

func add() {

	for i := 0; i < 5000; i++ {
		lock.Lock() // 加互斥锁
		x = x + 1
		lock.Unlock() // 解互斥锁
	}

	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()

	fmt.Println(x)
}
