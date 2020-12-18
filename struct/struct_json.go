package main

import (
	"encoding/json"
	"fmt"
)

// 1, 序列化： 把Go语言中的结构体变量 --> json 格式的字符串
// 2, 反序列化： json 格式的字符串 --> Go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age" db:"name" ini:"name"`
}

// 首先你需要知道的是，如果你能够为某个类型实现了MarshalJSON()([]byte, error)和UnmarshalJSON(b []byte) error方法，
// 那么这个类型在序列化（MarshalJSON）/反序列化（UnmarshalJSON）时就会使用你定制的相应方法。

// 序列化
//func (p *person) MarshalJSON(b []byte) error {
//
//	fmt.Printf("%#v\n", string(b))
//
//	return nil
//}

// 反序列化
//func (p *person) UnmarshalJSON(data []byte) (err error) {
//	required := struct {
//		Name string `json:"name"`
//		Age  string `json:"age"`
//	}{}
//	fmt.Printf("%#v\n", required)
//	err = json.Unmarshal(data, &required)
//
//	if err != nil {
//		return
//	}
//
//	return nil
//}

func main() {
	p1 := person{
		Name: "law",
		Age:  40,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("err:%v \n", err)
		return
	}
	fmt.Printf("%#v , %s \n", b, string(b))

	//fmt.Println(string(b))
	fmt.Printf("%v \n", string(b))

	// 反序列化
	str := `{"name":"xxiao", "age":19}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了在函数内部修改p2的值
	fmt.Printf("%#v \n", p2)

}
