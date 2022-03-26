package main

import "fmt"

// 反转字符串
func main() {
	tst := []byte{'H', 'E', 'L', 'L', 'O'}
	bytes := revers(tst)

	fmt.Printf("resutl: %#v", string(bytes))
}

func revers(s []byte) []byte {
	if len(s) <= 1 {
		return s
	}

	right := len(s) - 1
	left := 0

	for left < right {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}

	return s
}
