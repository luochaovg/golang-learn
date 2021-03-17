package main

import "fmt"

//交换律：a ^ b ^ c <=> a ^ c ^ b
//
//任何数于0异或为任何数 0 ^ n => n
//
//相同的数异或为0: n ^ n => 0
//
//var a = [2,3,2,4,4]
//
//2 ^ 3 ^ 2 ^ 4 ^ 4等价于 2 ^ 2 ^ 4 ^ 4 ^ 3 => 0 ^ 0 ^3 => 3

func main() {
	//a := 0
	//b := 6
	//
	//c := a ^ b
	//fmt.Println(c)

	num := []int{6, 1, 2, 3, 1, 2, 3}
	fmt.Println(singleNumber(num))
}

func singleNumber(nums []int) int {
	res := 0

	for _, v := range nums {
		res = res ^ v
	}

	return res
}
