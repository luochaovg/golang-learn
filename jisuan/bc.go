package jisuan

import "fmt"

// 包中标示符 变量名/函数名/结构体/接口等 如果首字母是小写，表示私有的，只能在这个包中使用
// 首字母大写的标示符可以被外部调用
func add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x + y
}

func init() {
	fmt.Printf("import 我时自动执行。。。")
}
