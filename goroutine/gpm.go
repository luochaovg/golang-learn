package main

import (
	"fmt"
	"runtime"
	"time"
)

func task(n int) {
	for {
		fmt.Println("这是工作", n)
		time.Sleep(time.Second)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	go task(1)
	go task(2)
	go task(3)
	go task(4)

	select {}
}
