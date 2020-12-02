package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := int32(97)
	fmt.Println(i)
	ret2 := string(i) //字符 ASII码
	fmt.Println(ret2)

	ret3 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v \n", ret3)

	// 从字符串中解析出对应的数据
	str := "10000"
	ret4, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", ret4)

	ret5, _ := strconv.Atoi(str) // string -> int
	fmt.Printf("%#v\n", ret5)

	str2 := strconv.Itoa(ret5) //int -> string
	fmt.Printf("%#v \n", str2)

	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr) // string -> bool
	fmt.Printf("%#v , %T \n", boolValue, boolValue)

	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64) // string -> float
	fmt.Printf("%#v , %T \n", floatValue, floatValue)

}
