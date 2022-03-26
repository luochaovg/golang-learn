package main

import (
	"fmt"
	"time"
)

// 参考连接： https://www.toutiao.com/i6995427307041079821

func main() {
	channelrange()
}

func arraydemo() {
	var array = [5]int{1, 2, 3, 4, 5}
	var array_new [5]int

	fmt.Println("原始数组的值:", array)
	fmt.Printf("原始数组的内存地址 %p: \n", &array)
	// range 后的操作对象 ， 是一个数据拷贝，而不是原始类型上操作
	// 值类型（string , array） 就是值拷贝
	// 引用类型 （slice map channel ）  就是引用拷贝
	for index, value := range array {
		if index == 0 {
			array[1] = 10 * value // 此array 跟range 后已经不是一个array 了
			array[2] = 20 * value
		}

		array_new[index] = value
	}
	fmt.Println("原始数据改变后的值", array)
	fmt.Printf("原始数据改变后的内存地址 %p: \n", &array)
	fmt.Println("依托原始数组创建的新数组的值", array_new)
}

func arrayslice() {
	var slice_src = []int{1, 2, 3, 4, 5}
	var slice_des = make([]int, 5)

	fmt.Println("原始切片的值：", slice_src)

	for index, value := range slice_src {
		if index == 0 {
			slice_src[1] = 10 * value // 此slice_src 跟range 后边的slice_src 指向同一内存，浅拷贝，引用复制
			slice_src[2] = 20 * value

			slice_src = append(slice_src, 6, 7, 8) // range 后边的slice_range 也是一个拷贝
		}

		slice_des[index] = value
	}

	fmt.Println("原始切片改变后的值", slice_src)      // 1，10，20，4，5，6，7，8
	fmt.Println("依托原始切片创建的新切片的值", slice_des) // 1，10，20，4，5
}

// range通过操作:=创建变量是一次性的
// range跟for搭配，是一个循环操作。但是操作符:=前边的变量确不是每次初始化，它只在第一次的时候分配内存，后边不论循环多少次都是修改该内存的值。
func channelrange() {
	var array = [5]int{1, 2, 3, 4, 5}
	var cha = make(chan *int, 5)
	var chb = make(chan *int, 5)

	for _, v := range array {
		cha <- &v

		tmp := v
		chb <- &tmp
	}

	close(cha)
	close(chb)

	//for {
	//	v, ok := <-cha
	//	if !ok {
	//		break
	//	}
	//	fmt.Println("cha", *v)
	//}
	time.Sleep(1 * time.Second)
	for v := range cha {
		fmt.Println("cha", *v)
	}

	fmt.Println("--------------------------")
	for v := range chb {
		fmt.Println("chb:", *v) // 输出1，2，3，4，5
	}
}
