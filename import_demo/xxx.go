package main

import (
	"fmt"

	calc "liwenzhou/jisuan"
)

// 包的路径从 GOPATH/src 后面的路径开始写起，路径分隔符用 /
// 想被别的包调用的标示符都要首字母大写
// 单行倒入和多行倒入
// 导入包的时候可以指定别名
// 导入包不想使用包内部的标示符， 需要使用匿名导入 '_'
// 每个包导入的时候会自动执行名为 init() 的函数, 他没有参数也没有返回值，也不能手动调用，多个包
// 	中都定义了init()函数，则他们的执行顺序 "相反"

//问题
//我在导入gopath目录下的包时报错“package xxx is not in GOROOT“，编译器没有去gopath下找包，
// 查了一下原因是GO111MODULE没有关， gomod 和 gopath 两个包管理方案，并且相互不兼容，
// 在 gopath 查找包，按照 goroot 和多 gopath 目录下 src/xxx 依次查找。
// 在 gomod 下查找包，解析 go.mod 文件查找包，mod 包名就是包的前缀，里面的目录就后续路径了。
// 在 gomod 模式下，查找包就不会去 gopath 查找，只是 gomod 包缓存在 gopath/pkg/mod 里面。
//
//解决方法
//把GO111MODULE置为off就行了。
//
//go env -w GO111MODULE=off

var x int = 5

// 全局声明 -> init() -> main()
func init() { // 没有参数也没有返回值，不能手动调用，程序自动调用
	fmt.Println("自动执行", x)
}

func main() {
	fmt.Println("aaa")
	ret := calc.Sub(10, 5)
	//ret := jisuan.Sub(10, 4)
	fmt.Println(ret)
}
