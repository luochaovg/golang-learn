package main

import (
	"container/list"
	"fmt"
)

// Binary Tree
type BinaryTree struct {
	Data  interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

// Constructor
func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{Data: data}
}

// 先序遍历-非递归
func (bt *BinaryTree) PreOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			res = append(res, t.Data) //visit
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 中序遍历-非递归
func (bt *BinaryTree) InOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			res = append(res, t.Data) //visit
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 后序遍历-非递归
func (bt *BinaryTree) PostOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	var preVisited *BinaryTree
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		v := stack.Back()
		top := v.Value.(*BinaryTree)
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right {
			res = append(res, top.Data) //visit
			preVisited = top
			stack.Remove(v)
		} else {
			t = top.Right
		}
	}
	return res
}
func main() {
	//t := NewBinaryTree(1)
	//t.Left = NewBinaryTree(3)
	//t.Right = NewBinaryTree(6)
	//t.Left.Left = NewBinaryTree(4)
	//t.Left.Right = NewBinaryTree(5)
	//t.Left.Left.Left = NewBinaryTree(7)
	//fmt.Println(t.PreOrderNoRecursion())
	//fmt.Println(t.InOrderNoRecursion())
	//fmt.Println(t.PostOrderNoRecursion())

	t := NewBinaryTree("A")
	t.Left = NewBinaryTree("B")
	t.Right = NewBinaryTree("C")

	t.Left.Left = NewBinaryTree("D")
	t.Left.Left.Left = NewBinaryTree("G")
	t.Left.Left.Right = NewBinaryTree("H")

	t.Right.Left = NewBinaryTree("E")
	t.Right.Right = NewBinaryTree("F")
	t.Right.Left.Right = NewBinaryTree("I")

	fmt.Println(t.PreOrderNoRecursion())
	//fmt.Println(t.InOrderNoRecursion())
	//fmt.Println(t.PostOrderNoRecursion())

}
