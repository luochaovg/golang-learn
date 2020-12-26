package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTimeTick32() int32 {
	return int32(time.Now().Unix())
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

// 基础做法 日期20191025时间戳1571987125435+3位随机数
func GenerateCode() {
	date := GetFormatTime(time.Now())
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(1000)
	a := GetTimeTick32()
	b := GetTimeTick64()
	fmt.Println(a, b)
	code := fmt.Sprintf("%s%d%03d", date, GetTimeTick32(), r)
	fmt.Println(r)
	fmt.Println(code, " rand ID generate successed!\n", len(code))
}

func main() {
	//todo 随机数可以用redis中的计数器代替 每天清0  每次取的时候先incr 分布式也是同理
	GenerateCode()
}
