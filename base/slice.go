package main

import "fmt"

func main() {
	var a []string
	var b []string
	a = append(a, "afafa")
	fmt.Println(a, len(a))
	fmt.Println(b == nil) // nil 没有内存空间

	// 1,初始化
	s1 := []int{1, 3, 64, 23}
	fmt.Println(s1, len(s1), cap(s1))

	// 2. 有数组得到切片
	c1 := [...]int{3, 5, 6, 7}
	c2 := c1[2:]
	c3 := c1[:2]
	c4 := c1[1:3] // 左闭右开

	c1[1] = 9000 // 修改底层数组指，切片指改变
	// 切片指向底层的一个数组， 切片是一个引用类型
	// 切片的长度就是它元素的个数
	// 容量： 底层数组切片第一个元素到最后一个元素的数量
	fmt.Println(c1, c2, c3, c4, cap(c3))

	// make 函数创建切片
	d1 := make([]int, 2, 8)
	fmt.Println(d1)

	// 切片就是一个框，框住了一块连续的内存， 连续的相同类型的内存地址空间
	// 切片属于引用类型， 真正的数据都是保存在底层数组里，
	// 切片不能直接比较 ，切片唯一合法的比较是和nil比较
	// 一个nil指的切片没有底层数组，一个nil值的切片的长度和容量都是0
	// 但是不能说 一个长度和容量都是0的切片一定是nil
	var f1 []string
	f2 := []string{}
	f3 := make([]string, 2, 8)
	fmt.Println(f1 == nil, len(f1), cap(f1))
	fmt.Println(f2 == nil, len(f2), cap(f2))
	fmt.Println(f3 == nil, len(f3), cap(f3))

	//  判断切片是否为空一定要使用
	fmt.Println(len(f3))

	f4 := []int{12, 3, 45, 4, 54, 54, 5}
	f5 := f4
	fmt.Println(f4, f5)

	f4[3] = 9999
	fmt.Println(f4, &f4[0])

	// append 追加元素，原来的底层数组放不下的时候，go就会把底层数组换一个内存地址空间
	// 必须用变量接收append 的返回值
	f4 = append(f4, 34)
	fmt.Println(f4, &f4[0])

	f8 := []int{111, 222, 333, 444}
	f9 := append(f4, f8...) // ... 把 f8 拆开

	fmt.Println(f8, f9, len(f9), cap(f9))

	//	copy
	f10 := make([]int, 10)
	copy(f10, f4)
	fmt.Println(f10, f4)

	// 切片不保存具体的值
	// 切片对应一个底层数组
	// 底层数组都是占用一块连续的内存
	x1 := [...]int{1, 3, 5} // 数组
	x2 := x1[:]             // 切片

	x2 = append(x2[:1], x2[2:]...)
	fmt.Println(len(x2), cap(x2))
	//fmt.Println(&x2[0], &x1[0])
	fmt.Println(*&x2[0], &x1[0])
	fmt.Println(&x2[1], &x1[1])
	x1[1] = 2324
	fmt.Println(x1)

	var aa = make([]int, 5, 10)
	fmt.Println(aa)
	for i := 0; i < 10; i++ {
		aa = append(aa, i)
	}

	fmt.Println(aa)

	urls := make(map[string]string, 3)
	// 这里随便个例子
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"

	names := make([]string, len(urls)) // 注意此时切片已经有值了 []string{ "", "", ""}
	//var names []string
	for key, _ := range urls {
		names = append(names, key) // append 会自动给切片执行初始化操作
	}

	fmt.Println(names, len(names))

}
