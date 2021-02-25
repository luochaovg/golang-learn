package main

import "fmt"

// 如何判断一个链表有没有闭环
// 需要将遍历过的节点存入map，以址为key，空struct为值
// 遍历时，当前节点是否已存在，存在即有环。
//https://leetcode-cn.com/problems/linked-list-cycle/submissions/

type LinkList struct {
	Val  int
	Next *LinkList
}

func HasRang(l *LinkList) bool {
	// 利用快慢指针
	slow := l
	fast := l

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}

func main() {

	node1 := new(LinkList)
	node2 := new(LinkList)
	node3 := new(LinkList)
	node4 := new(LinkList)
	node5 := new(LinkList)
	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node1

	b := HasRang(node1)
	fmt.Println(b)

}

// 走n个台阶，每次只能走一步或2步， 多少种走法
func f(n int) int {
	if n <= 2 {
		return n
	}

	return f(n-1) + f(n-2)
}
