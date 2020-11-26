package main

import "fmt"

func main() {
	//trace.Start(os.Stderr)
	//defer  trace.Stop()

	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	x, y, z := <-ch, <-ch, <-ch
	fmt.Println(x, y, z)
}
