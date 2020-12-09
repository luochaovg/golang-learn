package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
// GO语言内置了一些针对内置的基本数据类型的一些并发安全的操作
// atomic.AddInt64

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	//x++
	//lock.Lock()
	//x += 1
	//lock.Unlock()

	atomic.AddInt64(&x, 1)
	wg.Done()
}
func main() {
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go add()
	}

	wg.Wait()
	fmt.Println(x)

	// 比较并交换
	ok := atomic.CompareAndSwapInt64(&x, 100, 200)
	fmt.Println(ok, x)
}
