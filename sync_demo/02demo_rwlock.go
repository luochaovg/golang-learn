package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

// 读操作
func read() {
	defer wg.Done()

	rwlock.RLock() // 加读锁
	fmt.Println(x)
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
}

// 写
func write() {
	defer wg.Done()

	rwlock.Lock() // 加写锁
	x += 1
	time.Sleep(time.Millisecond * 5) // 假设读操作耗时5毫秒
	rwlock.Unlock()                  // 解写锁
}

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	//time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
