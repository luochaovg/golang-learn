package main

import (
	"fmt"
	"math"
	"strconv"
)

func main()  {
		//s1 := "34234.342434"
		s1 := "234.2349999"
		s2 := "1"

		//fmt.Println(StringAdd(s1, s2))
		fmt.Println(StringDivision(s1, s2))
}

func FormatFloat(num float64, decimal int) string  {
	d := float64(1)
	if decimal >0 {
		d = math.Pow10(decimal)
	}

	return strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
}


// StringAdd 字符串相加
func StringAdd(s1, s2 string) string {
	f1, err1 := strconv.ParseFloat(s1, 64)
	f2, err2 := strconv.ParseFloat(s2, 64)

	// 两个字符串都不是数字,返回无数据
	if err1 != nil && err2 != nil {
		return ""
	}

	return FormatFloat(f1+f2, 6)
	//return fmt.Sprintf("%.4f", f1+f2)
}

// StringSub 字符串相减
func StringSub(s1, s2 string) string {
	f1, err1 := strconv.ParseFloat(s1, 64)
	f2, err2 := strconv.ParseFloat(s2, 64)

	// 两个字符串都不是数字,返回无数据
	if err1 != nil && err2 != nil {
		return ""
	}
	return fmt.Sprintf("%.4f", f1-f2)
}

// StringDivision 字符串相除
func StringDivision(s1, s2 string) string {
	f1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		return ""
	}
	f2, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		return ""
	}
	// 分母不能为0
	if f2 == 0 {
		return ""
	}
	return FormatFloat(f1/f2, 6)
	//return fmt.Sprintf("%.4f", f1/f2)
}