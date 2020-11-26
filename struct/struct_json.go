package _struct

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

	//fmt.Println(string(b))
	fmt.Printf("%v \n", string(b))

	// 反序列化
	str := `{"name":"xxiao", "age":19}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了在函数内部修改p2的值
	fmt.Printf("%#v \n", p2)

}
