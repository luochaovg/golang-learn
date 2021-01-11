package main

import (
	"fmt"
	"time"
)

func time01() {
	now := time.Now()
	fmt.Println(now.Year(), now.Day())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	//fmt.Printf("%#v", now)
	//fmt.Printf("%v", now)

	//ret := time.Unix(1606457598, 0)
	//fmt.Println(ret.Year())
	//fmt.Println(time.Second)

	//fmt.Println(now.Add(time.Hour * 24))

	d := now.Sub(now.Add(time.Hour * 24 * 1))
	fmt.Println(d.Hours())

	//fmt.Println(222222)
	//n := time.Duration(5 * time.Second) // 时间间隔
	//time.Sleep(time.Second * 5)
	//fmt.Println(33333)
	// 定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t) // 1秒执行一次
	//}

	//  2006/1/2 15:04:05
	// Y     m  d  H  M  S
	// 2006  1  2  3  4  5

	// 格式化时间
	//fmt.Println(now.Format("2006-01-02"))
	//fmt.Println(now.Format("2006/01/02 15:04:05.000 PM"))

	// 把一个字符串的时间转换成时间戳
	//timeObj, err := time.Parse("2006-01-02", "1998-09-23")
	//if err != nil {
	//	fmt.Printf("err %v\n", err)
	//}
	//fmt.Println(timeObj.Unix())
}

func main() {
	now := time.Now() // 本地时间
	fmt.Println(222222)
	fmt.Println(now)

	// 明天的这个时间
	timea, _ := time.Parse("20060102", "20201128")

	fmt.Println("aaaaaa", timea.Format("2006-01-02"))

	// 按照东八区格式解析时间
	l, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("11, %v, \n", err)
	}

	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-11-28 14:31:50", l)
	if err != nil {
		fmt.Printf("22, %v, \n", err)
	}

	fmt.Println(timeObj, 444)

	td := timeObj.Sub(now)
	fmt.Println(td)

}
