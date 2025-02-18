package main

import "fmt"

func demoSl(s []int) {

	s[0] = 12
	fmt.Printf("demos2 sl:%#v", s)
}

// 复习： https://www.jianshu.com/p/ae8a413fc33f
func main() {
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1:%#v", s1)
	//demoSl(s1)
	deomS3(s1)

	fmt.Printf("s2:%#v", s1)

	//3,创建一个 slice := make([]int, 5, 10), 然后 slice[8] 和 slice[:8] 的运行结果是什么?
	//数据直接访问(slice[index])时, index 值不能超过 len(slice) 范围
	//创建切片(slice[start:end])时, start 和 end 指定的区间不能超过 cap(slice) 范围
	//slice[8] 会 panic, 而 slice[:8] 正常返回.

}

func deomS3(s []int) {
	s[0] = 3

	fmt.Printf("s:%#v", s)
}

func demoS2() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]                   // bc
	str2[1] = "new"                    // str2 = bnew
	fmt.Println(str1)                  //  str1 = abnew
	str2 = append(str2, "z", "x", "y") // str2 = bcnewzxy
	fmt.Println(str1)                  // str1 = abnew
}
