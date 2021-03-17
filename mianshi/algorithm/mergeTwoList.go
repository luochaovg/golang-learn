package main

import "fmt"

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
https://leetcode-cn.com/problems/merge-two-sorted-lists/
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func mergeTwoLists(l1 *ListNode3, l2 *ListNode3) *ListNode3 {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var res *ListNode3
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)

	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}

	return res
}

func main() {
	l1 := &ListNode3{
		Val: 1,
		Next: &ListNode3{
			Val: 2,
			Next: &ListNode3{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode3{
		Val: 1,
		Next: &ListNode3{
			Val: 3,
			Next: &ListNode3{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := mergeTwoLists(l1, l2)

	fmt.Println(res)
}
