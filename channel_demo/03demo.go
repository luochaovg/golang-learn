package main

import "fmt"

// 关闭通道
// 当通道被关闭时，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值，
// 再然后取到的值一直都是对应类型的零值。那如何判断一个通道是否被关闭了呢？
func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20

	close(ch1)

	//for x := range ch1 {
	//	fmt.Println(x)
	//}

	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok)
}
