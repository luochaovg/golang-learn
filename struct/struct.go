package _struct

import "fmt"

// 自定义类型和类型别名
// type 后面跟的是类型
type myInt int    // 自定义类型
type youInt = int // 类型别名

// struct 关键字， 表示结构体
// 结构体 是值类型的， 赋值的时候就是copy
type person struct {
	name   string
	age    int
	hobby  []string
	gender string
}

type girl struct {
	h, w int
}

// go 语言函数传参数永远是拷贝
// 函数内改参数值，传入指针类型
func f(g *girl) {
	(*g).h = 34 // g.h = 34
}

func main() {

	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T \n", n)

	var m youInt
	m = 200
	fmt.Printf("%T \n", m)

	var c rune // int32 字符
	c = '中'
	fmt.Printf("%T \n", c)

	var law person
	law.name = "luochao"
	law.age = 33
	law.gender = "男"
	//law.hobby = []string{"篮球", "足球"}

	fmt.Printf("%T \n", law)
	fmt.Println(law, law.gender)

	//声明s变量， 为匿名结构体, 多用于一些临时场景
	var s struct {
		x string
		y int
	}
	s.x = "fjadf"
	s.y = 100
	fmt.Printf("%T value:%v \n", s, s)

	var g girl
	g.h = 150
	g.w = 50
	f(&g)

	fmt.Printf("%T value:%v \n", g, g)

	// 结构体指针 1
	// new 申请内存使用， 特定的类型， int,string,bool, struct, 返回指针类型，值为内存地址
	// make 申请内存使用， slice , map , channel , 返回的是类型
	var p2 = new(person)
	(*p2).name = "llll"
	p2.gender = "男"

	fmt.Printf("%T\n", p2)
	fmt.Printf(" %p \n", p2)  // p2 保存的值是一个内存地址
	fmt.Printf(" %p \n", &p2) // p2的内存地址

	// 结构体指针 2.1 - key - value 初始化 ， 推荐
	var p3 = &person{
		name:   "xx",
		gender: "女",
	}
	fmt.Printf("%T , %#v , %p\n", p3, p3, p3)

	// 结构体指针 2.2 - 使用值列表的形式,值的顺序和结构体定义时字段的顺序一致
	p4 := person{
		"lmy",
		0,
		[]string{"篮球"},
		"女",
	}
	fmt.Println(p4)

	p5 := newPerson("law", 23)
	fmt.Println(p5)

	d1 := newDog("zhoulin")
	d1.wang()

	p5.gn()
	fmt.Println(p5.age)
	p5.gn2()
	fmt.Println(p5.age)
	p5.like()
	fmt.Printf("%T\n", p5)

	mm := myInt(100)
	mm.intHello()

}

// 构造函数： 约定俗成用 new 开头
// 构造函数， 返回一个结构体变量的函数
// 返回的结构体还是结构体指针 ， 当结构体比较大的时间尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 标示符：变量名， 函数名， 类型名
// go 语言中如果标示符首字母大写，表示对外部可见（公共函数）， 大写的需要写注释
type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{name: name}
}

// 方法是作用与特定类型的函数
// 接收者表示的是调用该方法的具体类型变量， 多用类型名首字母小写表示,
func (d dog) wang() {
	fmt.Printf("%s wa wa wa \n", d.name)
}

// 指针接收者
// 1.需要修改结构体变量时要使用指针接收者
// 2.结构体比较大
// 3.保持一致性
func (d *dog) wang1() {
	d.name = "aaa"
	fmt.Printf("%s wa wa wa \n", d.name)
}

// 使用值接收者 （传copy进去）
func (p person) gn() {
	p.age += 1
}

// 指针接收者 （传内存地址）(推荐使用)
func (p *person) gn2() {
	p.age += 1
}

func (p *person) like() {
	p.hobby = []string{"篮球"}
}

// 接收的方法不能是内置的，可以自己造一个
// 不能给别的包里面的类型添加方法， 只能给自己的包里的类型添加方法
func (m myInt) intHello() {
	fmt.Println("我是一个int")
}
