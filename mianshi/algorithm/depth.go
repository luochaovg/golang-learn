package main

import "fmt"

/**
二叉树的最大深度
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
*/

func main() {
	t := NewTreeNode2(3)
	t.Left = NewTreeNode2(5)
	t.Right = NewTreeNode2(1)
	t.Left.Left = NewTreeNode2(6)
	t.Left.Right = NewTreeNode2(2)
	t.Left.Right.Left = NewTreeNode2(7)
	t.Left.Right.Right = NewTreeNode2(4)
	t.Right.Left = NewTreeNode2(0)
	t.Right.Right = NewTreeNode2(8)
	t.Right.Right.Right = NewTreeNode2(4)
	t.Right.Right.Right.Right = NewTreeNode2(8)

	fmt.Println(maxDepth(t))
	fmt.Println(minDept(t))
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func NewTreeNode2(v int) *TreeNode2 {
	return &TreeNode2{
		Val: v,
	}
}

func maxDepth(root *TreeNode2) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left < right {
		return right + 1
	} else {
		return left + 1
	}
}

func minDept(root *TreeNode2) int {
	if root == nil {
		return 0
	}

	l := minDept(root.Left)
	r := minDept(root.Right)

	if l == 0 || r == 0 {
		return l + r + 1
	}

	if r < l {
		return r + 1
	}

	return l + 1
}
