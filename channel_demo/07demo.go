package main

import (
	"fmt"
	"sync"
)

// 空 struct 控制channel
var wg sync.WaitGroup

func work(ch chan struct{}) {
	defer wg.Done()

	fmt.Println("start work")

	//time.Sleep(time.Second * 2)
	ch <- struct{}{}
	fmt.Println("end work")
	close(ch)
}

func main() {
	ch := make(chan struct{})

	wg.Add(1)
	go work(ch)
	<-ch
	wg.Wait()
	fmt.Println("end")
}
