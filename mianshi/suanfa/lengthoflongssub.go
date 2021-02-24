package main

import "fmt"

//3. 无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	subMap := make(map[byte]int)
	for i, j := 0, 0; j < n; j++ {
		if v, ok := subMap[byte(s[j])]; ok {
			if i < v {
				i = v
			}
		}
		if ans < (j - i + 1) {
			ans = j - i + 1
		}
		subMap[byte(s[j])] = j + 1
	}
	return ans
}

func main() {
	s := "abcabcbb"
	n := lengthOfLongestSubstring(s)

	fmt.Println(n)
}
