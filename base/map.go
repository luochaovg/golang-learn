package main

import (
	"fmt"
	//"unsafe"
)

func main() {
	deomo()
	var m1 map[string]int

	fmt.Println(m1 == nil) // 还没有初始化（没有内存中开辟空间  ）

	m1 = make(map[string]int, 3) // 要估算好改map的容量， 避免在程序允许期间在动态扩容
	// panic: assignment to entry in nil map

	m1["age"] = 49
	m1["height"] = 173
	m1["height2"] = 1733
	fmt.Printf("m1:%#v \n", m1)
	resetmap(m1)
	fmt.Printf("m1:%#v \n", m1)

	return

	fmt.Println(m1)
	fmt.Println(m1["age"], m1["abc"]) // 49 0

	v, ok := m1["nihao"] // ok 一定是bool类型， 约定成俗用ok接收， 用其他字符也可以
	if !ok {
		fmt.Println("no key")
	} else {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 删除的key不存在什么都不干
	// 存在key， 正常删除
	delete(m1, "age")
	fmt.Println(m1)

	//  map 是无序的， hash 类型

	// map 和 slice 的组合
	// 元素为map的切片
	s1 := make([]map[int]string, 2, 10)

	//s1[0][200] = "A" // 没有对内部map初始化

	s1[0] = make(map[int]string, 1)
	s1[0][1] = "AAA"
	fmt.Println(s1)

	// 值为切片类型的map
	var m2 = make(map[int][]string, 10)
	m2[12] = []string{"a", "b"} // 初始化

	fmt.Println(m2)

}

func resetmap(m map[string]int) {
	m["age"] = 50344
}

type People2 struct {
	Name string
}

func deomo() {
	m := map[string]People2{"name": {Name: "xx"}}
	delete(m, "name")

	//golang里面的map是通过hashtable来实现的，具体方式就是通过拉链法（数组+链表）来实现的，这里对比c++的map,c++里面的map， 是通过红黑树来实现的。
	//所以二者在遍历的时候做删除操作，golang的是可以直接操作的，因为内部实现是哈希映射，删除并不影响其他项，
	//而c++中的map删除，由于是红黑树，删除任意一项，都会打乱迭代指针，不能再O(1)时间内删除。
	//同时，golang里面的key是无序的，即使你顺序添加，遍历的时候也是无序。
	//golang里面的map,当通过key获取到value时，这个value是不可寻址的，因为map 会进行动态扩容，当进行扩展后，map的value就会进行内存迁移，
	//其地址发生变化，所以无法对这个value进行寻址。也就是造成上述问题的原因所在。map的扩容与slice不同，那么map本身是引用类型，作为形参或返回参数的时候，
	//传递的是值的拷贝，而值是地址，扩容时也不会改变这个地址。而slice的扩容，会导致地址的变化。

	//MAP的原理
	//type hmap struct {
	//	count        int  //元素个数
	//	flags        uint8
	//	B            uint8 //扩容常量
	//	noverflow    uint16 //溢出 bucket 个数
	//	hash0        uint32 //hash 种子
	//	buckets      unsafe.Pointer //bucket 数组指针
	//	oldbuckets   unsafe.Pointer //扩容时旧的buckets 数组指针
	//	nevacuate    uintptr  //扩容搬迁进度
	//	extra        *mapextra //记录溢出相关
	//}
	//
	//type bmap struct {
	//	tophash        [bucketCnt]uint8
	//	// Followed by bucketCnt keys
	//	//and then bucketan Cnt values
	//	// Followed by overflow pointer.
	//}
	//
	//每个map的底层结构是hmap，是有若干个结构为bmap的bucket组成的数组，每个bucket可以存放若干个元素(通常是8个)，
	//那么每个key会根据hash算法归到同一个bucket中，当一个bucket中的元素超过8个的时候，hmap会使用extra中的overflow来扩展存储key。

}
