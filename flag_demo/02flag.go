package main

import (
	"flag"
	"fmt"
	"time"
)

// flag获取命令行参数
func main() {
	// 创建一个标识位参数
	name := flag.String("name", "default name", "请输入名字") // 得到的指针
	age := flag.Int("age", 9000, "请输入真实年龄")
	married := flag.Bool("married", false, "请输入结婚了吗")
	mTime := flag.Duration("ct", time.Second, "结婚多久了")

	// 使用flag
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*mTime)

	//var name string
	//flag.StringVar(&name, "name", "law", "input you name")
	//// 解析变量值
	//flag.Parse()
	//fmt.Println(name)

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

}
