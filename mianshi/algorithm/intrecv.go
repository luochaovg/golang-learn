package main

import (
	"fmt"
	"math"
)

// 整数反转
//给你一个 32 位的有符号整数 x ，返回 x 中每位上的数字反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	fmt.Println(math.MaxInt32)
	num := 0

	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}

	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}

	return num
}

func main() {
	var x int = 1534236469

	m := reverse(x)
	fmt.Println(m)
}
