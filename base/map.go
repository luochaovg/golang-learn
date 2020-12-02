package base

import "fmt"

func main() {

	var m1 map[string]int

	fmt.Println(m1 == nil) // 还没有初始化（没有内存中开辟空间  ）

	m1 = make(map[string]int, 3) // 要估算好改map的容量， 避免在程序允许期间在动态扩容
	// panic: assignment to entry in nil map

	m1["age"] = 49
	m1["height"] = 173

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
