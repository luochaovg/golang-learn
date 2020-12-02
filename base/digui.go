package base

import "fmt"

func jies(x uint64) uint64 {
	if x <= 1 {
		return 1
	}
	return x * jies(x-1)
}

func taijie(x uint64) uint64 {

	if x == 1 {
		return 1
	} else if x == 2 {
		return 2
	}

	return taijie(x-1) + taijie(x-2)
}

// 递归斐波那契
func fbo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return fbo1(n-1) + fbo1(n-2)
	} else {
		return -1
	}
}

func main() {
	//fmt.Println(jies(5))
	fmt.Println(taijie(4))
}
