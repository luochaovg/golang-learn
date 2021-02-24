package main

import "fmt"

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
链接：https://leetcode-cn.com/problems/add-two-numbers

 * Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	jw := sum / 10
	sum = sum % 10

	head := &ListNode{sum, nil}
	tmp := head

	val1, val2 := 0, 0
	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil || l2 != nil || jw != 0 {
		if l1 == nil {
			val1 = 0
		} else {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			val2 = 0
		} else {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum = val1 + val2 + jw
		jw = sum / 10
		sum = sum % 10
		tmp.Next = &ListNode{sum, nil}
		tmp = tmp.Next

	}

	return head

}

func main() {

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := addTwoNumbers(l1, l2)
	fmt.Printf("%v \n", l3)
	fmt.Printf("%v \n", l3.Next)
}
