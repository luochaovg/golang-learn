package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	// for range 循环的时候会创建每个元素的副本，而不是元素的引用，
	// 所以 m[key] = &val 取的都是变量 val 的地址，
	// 所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.
	for key, val := range slice {
		fmt.Printf("v :  %#v \n", &val)
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}

}

// 正确的写法
func is_true() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "===>", *v)
	}
}
