package main

import "fmt"

// https://mp.weixin.qq.com/s/kEQI74ge6VhvNEr1d3JW-Q
// 内存操作 https://www.flysnow.org/2017/07/06/go-in-action-unsafe-pointer.html
func main() {
	//strSlice := []string
	strSlice := make([]string, 5, 5)

	strSlice[0] = "lc"

	fmt.Printf("one: %#v \n", strSlice)

	testS(strSlice)
	fmt.Printf("two: %#v \n", strSlice)
}

// 某个函数的参数为切片时， 并没有复制拷贝， 隐式地址
func testS(params []string) {
	params = append(params, "xx") // 发生了地址内存拷贝
	//params[1] = "xx" // 没有发生拷贝，还是原来传入的数组
}
