package main

import (
	"fmt"
	"strconv"
	"sync"
)

// GO 内置的map 不是并发安全的

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			set(key, n)
//			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

var m2 = sync.Map{}

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 必须使用sync.Map 内置的Store 方法存储值
			value, _ := m2.Load(key) // 必须使用sync.Map 使用的Load 方法取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

	//fmt.Printf("%#v, \n ", m2)
}
