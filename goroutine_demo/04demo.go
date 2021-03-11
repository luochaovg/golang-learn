package main

import "fmt"

func Cat(ch chan int) {
	fmt.Println("cat")
	ch <- 0
}
func Dog(ch chan int) {
	fmt.Println("dog")
	ch <- 0
}
func Fish(ch chan int) {
	fmt.Println("fish")
	ch <- 0
}

func main() {
	ch := make(chan int, 1)

	for i := 0; i < 100; i++ {
		go Cat(ch)
		<-ch
		go Dog(ch)
		<-ch
		go Fish(ch)
		<-ch
	}
}
