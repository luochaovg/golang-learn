package split_string

import (
	"strings"
)

// 切割字符串

// example:
// abcbdfb , b
func Split(str string, sep string) []string {
	var ret = make([]string, 0, strings.Count(str, sep)+1)

	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}

	ret = append(ret, str)
	return ret
}

// 斐波那契函数
func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}
