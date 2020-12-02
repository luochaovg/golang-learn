package base

import (
	"fmt"
	"strings"
)

var name string = "luochao"
var age int = 21
var isBoy bool = true
var wift string

// 注： 非全局变量声明必须使用，全局变量可以不用
var (
	name1 string
	age2  int
)

// 批量声明常量是， 如果某一行声明后没有赋值，默认就和上一行一样
// iota 是常量计数器,, 在const 关键字出现时被重置为0，const 中没新增一行常量声明将使 iota 计数一次
const s1 = "abc"
const (
	s3     = 23
	s4     = "abc"
	s5     = 3
	s6     = iota
	s7     = iota
	s8, s9 = iota + 1, iota + 2
)

func main() {
	wift = "xiaoxue"

	fmt.Print(wift)
	fmt.Printf("name: %s \n", wift)
	fmt.Println("Hello word")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(name1)

	var s2 = "fasfas"
	s1 := "safa" // 简短变量声明， 只能在函数内用
	fmt.Println(s1, s2)

	// 匿名变量 下划线 "_"
	// 同一个作用域， 不能重复声明同名变量
	fmt.Println(s2, s3, s5, s6, s7, s8, s9)

	// 查看变量类型
	fmt.Printf("%T \n", s3)

	s8 := 12.23243434
	s9 := float32(12.23243434)
	fmt.Println(s8)
	fmt.Printf("%T \n", s8) // 默认Go 小数都是float64 类型
	fmt.Printf("%T \n", s9)

	// go 字符串使用双引号包裹的
	a1 := "afadfa"
	// go 单引号包裹的是字符 ， 单独的字母，汉字，符号表示一个字符
	a2 := 'h'

	fmt.Println(a1, a2)

	a3 := "hello wordfl"

	fmt.Println(a1+a3, len(a3))
	a4 := fmt.Sprintf("%s%s", a1, a3)
	fmt.Println(a4)

	// 分隔
	ret := strings.Split(a1, "a")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(a1, "a"))

	//strings.HasPrefix() 前缀
	//strings.HasSuffix() 后缀
	fmt.Println(strings.Index(a3, "w"))
	fmt.Println(strings.LastIndex(a3, "l"))
	//拼接
	fmt.Println(strings.Join(ret, "/"))

	a5 := "aabbccdd你好ndf"
	fmt.Println(len(a5)) // 字符串所有字节的长度， 并不是字符长度
	fmt.Println(a5[0])   // a 字符的 ascii 码
	fmt.Printf("%c\n", a5[0])

	//for i := 0; i < n; i++ {
	//	fmt.Println(a5[i])
	//}

	//for k, c := range a5 {
	//	fmt.Println(k, c)
	//}

	a7 := "白萝卜"
	a8 := []rune(a7) // 把字符串强制转换成一个rune 切片 ， int32
	fmt.Println(a8)
	a8[0] = '红'
	fmt.Println(string(a8)) // 把rune切片强制转成字符串

	a9 := byte('a')
	fmt.Println(a9)

	n1 := 12
	var f float64
	f = float64(n1)
	fmt.Println(f)

	fmt.Println("Aaaaaa")

	ss := "是反复看风景阿拉山口风景abcfadsfa"

	var m1 []string
	for _, v := range ss {
		//fmt.Println(v)
		//fmt.Printf("%c\n", v)
		//fmt.Printf("%T\n", v)
		//fmt.Println(string(v))
		m1 = append(m1, string(v))
	}

	fmt.Println(m1)
	fmt.Printf("%T", m1)

}
