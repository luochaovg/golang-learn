package main

// 如何判断一个链表有没有闭环

type a struct {
	val  int
	next *a
}

func main() {

}

// 走n个台阶，每次只能走一步或2步， 多少种走法
func f(n int) int {
	if n <= 2 {
		return n
	}

	return f(n-1) + f(n-2)
}
