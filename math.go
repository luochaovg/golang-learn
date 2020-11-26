package main

import "fmt"

func main() {

	a := 7
	a++

	fmt.Println(a)

	// & 按位与
	// 5 二进制 101
	// 3 二进制 011
	fmt.Println(5 & 3)
	fmt.Println(5 | 3)
	fmt.Println(5 << 1) // 1010  = 2*2^3 + 2*2^1
	fmt.Println(5 >> 1) // 10 = 2
}
