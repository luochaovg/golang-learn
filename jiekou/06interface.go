package main

import "fmt"

// 空接口 , 空接口没有必要起名字， 通常定义成下面的格式
// interface {}
// 所有的类型都实现了空接口， 也就是任意类型的变量都能保存到空接口中

// interface 是关键字
// interface{] 空接口类型

// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("%T, %v \n", a, a)

	// 类型段言
	//str, ok := a.(string)
	//if !ok {
	//	fmt.Println("猜错了")
	//} else {
	//	fmt.Println(str)
	//}

	//str, ok := a.(type)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串：", t)
	case int:
		fmt.Println("是一个int：", t)
	case int64:
		fmt.Println("是一个64：", t)

	case bool:
		fmt.Println("是一个bool：", t)
	}

}
func main() {

	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)

	m1["name"] = "luochao"
	m1["age"] = 900
	m1["like"] = [...]string{"篮球", "羽毛球"}
	m1["hobby"] = []string{"唱", "跳"}

	fmt.Printf("%T \n", m1)
	fmt.Println(m1)

	show(false)
	//show(nil)
	show("law")
	//show(m1)
}
