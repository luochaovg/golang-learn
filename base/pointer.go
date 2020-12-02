package base

import "fmt"

func main() {
	var a *int // nil
	fmt.Println(a)

	var a2 = new(int) // new 函数申请一个内存地址, string\int 返回的是对应类型的指针
	fmt.Println(a2)
	fmt.Println(*a2)

	*a2 = 100
	//
	fmt.Println(*a2)

	// make 也是用户内存分配 （slice , map, chan），make 函数返回的是对应的这三个类型本身

}
