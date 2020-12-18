package main

import "fmt"

// 神奇的指针
// 二个例子是使用 Go 语言经常会犯的错误1。当我们在遍历一个数组时，如果获取 range 返回变量的地址并保存到另一个数组或者哈希时，会遇到令人困惑的现象，下面的代码会输出 “3 3 3”：
// 一些有经验的开发者不经意也会犯这种错误，正确的做法应该是使用 &arr[i] 替代 &v，我们会在下面分析这一现象背后的原因。

func f1() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		fmt.Printf("%#v , %T\n", v, v)
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

func main() {
	f1()
}
