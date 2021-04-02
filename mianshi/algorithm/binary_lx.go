package main

import (
	"container/list"
)

type NodeBinary struct {
	Val   int
	Left  *NodeBinary
	Right *NodeBinary
}

func NewNodeBinary(num int) *NodeBinary {
	return &NodeBinary{
		Val: num,
	}
}

// 先序遍历 -- 非递归
func (bt *NodeBinary) XianXu() []int {
	t := bt
	stack := list.New()
	res := []int{}

	for t != nil || stack.Len() != 0 {
		if t != nil {
			res = append(res, t.Val)
			stack.PushBack(t)
			t = t.Left
		}

		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*NodeBinary)
			t = t.Right
			stack.Remove(v)
		}
	}

	return res
}

func main() {

}
