package main

import (
	"fmt"
	"sync"
)

// 参考 https://studygolang.com/articles/18042
// github ： https://github.com/chain-zhang/gomo/tree/master/dstr (  双向链表/ 单向链表 / 队列 / 栈 	)

/**
双向链表 Demo
*/
type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}

type DLink struct {
	mutex *sync.RWMutex
	Size  uint
	Head  *Node
	Tail  *Node
}

// 双链表初始化
func (d *DLink) Init() {
	d.mutex = new(sync.RWMutex)
	d.Size = 0
	d.Head = nil
	d.Tail = nil
}

// Get 获取指定位置的节点
func (d *DLink) Get(index uint) *Node {
	if d.Size == 0 || index > d.Size-1 {
		return nil
	}

	if index == 0 {
		return d.Head
	}

	node := d.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}

	return node
}

// 链表节点的新增分为两种，一种是在链表后面追加节点，该方式，我们称为append；另外一种方式是在指定位置插入节点，我们叫做insert。
// Append 向双链表后面追加节点
func (d *DLink) Append(node *Node) bool {
	if node == nil {
		return false
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.Size == 0 {
		d.Head = node
		d.Tail = node
		node.Next = nil
		node.Prev = nil
	} else {
		node.Prev = d.Tail
		node.Next = nil
		d.Tail.Next = node
		d.Tail = node
	}

	d.Size++
	return true
}

// Insert 向双链表指定位置插入节点
func (d *DLink) Insert(index uint, node *Node) bool {
	if index > d.Size || node == nil {
		return false
	}

	if index == d.Size {
		return d.Append(node)
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if index == 0 {
		node.Next = d.Head
		d.Head = node
		d.Head.Prev = nil
		d.Size++
		return true
	}

	nextNode := d.Get(index)
	node.Prev = nextNode.Prev
	node.Next = nextNode
	nextNode.Prev.Next = node
	nextNode.Prev = node
	d.Size++
	return true
}

// 删除指定位置节点
func (d *DLink) Delete(index uint) bool {
	if index > d.Size-1 {
		return false
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if index == 0 {
		if d.Size == 1 {
			d.Head = nil
			d.Tail = nil
		} else {
			d.Head.Next.Prev = nil
			d.Head = d.Head.Next
		}
		d.Size--
		return true
	}
	if index == d.Size-1 {
		d.Tail.Prev.Next = nil
		d.Tail = d.Tail.Prev
		d.Size--
		return true
	}

	node := d.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	d.Size--
	return true
}

// Display 打印双链表信息
func (d *DLink) Display() {
	if d == nil || d.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", d.Size)
	ptr := d.Head
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}

// Reverse 倒序打印双链表信息
func (d *DLink) Reverse() {
	if d == nil || d.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", d.Size)
	ptr := d.Tail
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Prev
	}
}
func main() {
	var test = new(DLink)
	test.Init()
	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2}
	node3 := &Node{Data: 3}
	test.Append(node1)
	test.Append(node2)
	test.Append(node3)

	test.Display()
	test.Reverse()
}
