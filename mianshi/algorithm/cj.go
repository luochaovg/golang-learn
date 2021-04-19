package main

import (
	"fmt"
	"sort"
)

//import "sort"

// lfu 算法
// https://leetcode-cn.com/problems/lfu-cache/

// 快排
// https://leetcode-cn.com/problems/sort-an-array/
func SortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := nums[0]
	var (
		left  []int
		right []int
	)

	for i := 1; i < len(nums); i++ {
		if middle < nums[i] {
			right = append(right, nums[i])
		} else {
			left = append(left, nums[i])
		}
	}

	left = SortArray(left)
	right = SortArray(right)

	left = append(left, middle)

	return append(left, right...)
}

// 归并 排序链表
// https://leetcode-cn.com/problems/sort-list/
// https://leetcode-cn.com/problems/sort-list/submissions/
// https://leetcode-cn.com/problems/sort-list/solution/sort-list-gui-bing-pai-xu-lian-biao-by-jyd/
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// 如果 head为空或者head就一位,直接返回
	if head == nil || head.Next == nil {
		return head
	}
	// 定义快慢俩指针,当快指针到末尾的时候,慢指针肯定在链表中间位置
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// 把链表拆分成两段,所以设置中间位置即慢指针的next为nil
	n := slow.Next
	slow.Next = nil
	// 递归排序
	return merge(sortList(head), sortList(n))
}

func merge(node1 *ListNode, node2 *ListNode) *ListNode {
	// 设置一个空链表,
	node := &ListNode{Val: 0}
	current := node
	// 挨个比较俩链表的值,把小的值放到新定义的链表里,排好序
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			current.Next, node1 = node1, node1.Next
		} else {
			current.Next, node2 = node2, node2.Next
		}
		current = current.Next
	}

	// 两链表可能有一个没走完,所以要把没走完的放到链表的后面
	// 注意,此处跟 数组不一样的是, 数组为什么要循环,因为数组可能一个数组全部走了(比如 12345与6789比较, 前面的全部走完,后面一个没走),另一个可能有多个没走..
	// 链表虽然也有这种可能,但是 node1和node2已经是有序的了,如果另外一个没有走完,直接把next指向node1或者node2就行,因为这是链表
	if node1 != nil {
		current.Next, node1 = node1, node1.Next
	}
	if node2 != nil {
		current.Next, node2 = node2, node2.Next
	}
	return node.Next
}

// 反转链表 K 个一组翻转链表
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// 路径综合
// https://leetcode-cn.com/problems/path-sum-ii/

//206. 反转链表 递归
//https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return p
}

// 逐个将旧链表的节点插入到新链表 迭代
func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode

	for head != nil {
		tmp := head
		head = head.Next

		tmp.Next = newHead
		newHead = tmp
	}

	return newHead
}

//15. 三数之和
// https://leetcode-cn.com/problems/3sum/
// https://leetcode-cn.com/problems/3sum/submissions/
func threeSum(nums []int) [][]int {
	// 排序，快排
	sort.Ints(nums)
	// 初始化数据
	result := [][]int{}
	len := len(nums)

	//outer:
	// left 表示第一个运算数索引
	for left := 0; left < len-2; left++ {
		//若最小值，比 0 大，则一定没有结果
		if nums[left] > 0 {
			break
		}

		// 不能出现重复数字，因此需要判断若值相同，不参与比较，比较哪些不相同的
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		// i, j 表示第 2， 3 个运算数索引
		// i 从 第一个运算数的下一个开始
		// j 从最后一个元素开始
		i, j := left+1, len-1
		// 加上最大的两个数，还不能大于 0
		if nums[left]+nums[j-1]+nums[j] < 0 {
			continue
		}
		// 比较目标
		target := 0 - nums[left]

		// 迭代比较
		for i < j {
			sum := nums[i] + nums[j]
			if sum == target {
				// 找到一组
				result = append(result, []int{nums[left], nums[i], nums[j]})
				//break outer
				// [-4, -1, -1, 0, 1, 2]
				// 准备找下一组，重置索引
				// 不能出现重复数字，因此需要判断若值相同，不参与比较，比较哪些不相同的
				for i++; i < j && nums[i] == nums[i-1]; i++ {
				}
				for j--; i < j && nums[j] == nums[j+1]; j-- {
				}
			} else if sum < target {
				// 小于，则将小的数增加，就是
				i++
			} else {
				// 大于, 则将小的数减小
				j--
			}
		}
	}

	return result

}

func main() {
	//nums := []int{2, 13, 42, 34, 56, 23, 67, 365, 87665, 54, 68, 3}

	//sort.Ints(nums)
	//fmt.Println(SortArray(nums))

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
	rl1 := reverseList2(l1)
	fmt.Printf("%#v , %#v", l1, rl1)

}
