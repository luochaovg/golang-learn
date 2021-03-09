package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val: v,
	}
}

func lowestCommonAncestorDemo(root, p, q *TreeNode) *TreeNode { // 236. 二叉树的最近公共祖先
	if root == nil { // 此时不可能查询到结果
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val { // 子树中寻找到结果节点，返回root
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)   // 寻找左子树
	right := lowestCommonAncestor(root.Right, p, q) // 寻找右子树

	if left == nil { // 从下一层返回来的查询结果为nil 没有找到
		return right
	} else if right == nil { // 从下一层返回来的查询结果为nil 没有找到
		return left
	} else { // 当左右子树都找到时返回root
		return root
	}

	return nil // 当在此棵子树上进行查找无pq时，返回nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}

	return nil
}

func main() {
	t := NewTreeNode(3)
	t.Left = NewTreeNode(5)
	t.Right = NewTreeNode(1)
	t.Left.Left = NewTreeNode(6)
	t.Left.Right = NewTreeNode(2)
	t.Left.Right.Left = NewTreeNode(7)
	t.Left.Right.Right = NewTreeNode(4)
	t.Right.Left = NewTreeNode(0)
	t.Right.Right = NewTreeNode(8)

	p := &TreeNode{
		Val: 5,
	}

	q := &TreeNode{
		Val: 4,
	}

	fmt.Println(lowestCommonAncestor(t, p, q))

}
