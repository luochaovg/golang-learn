package main

import "fmt"

func DeferFunc1(i int) (t int) { // 4
	t = i
	defer func() {
		t += 3
	}()
	return t
}

//创建变量t并赋值为1
//执行return语句，注意这里是将t赋值给返回值，此时返回值为1（这个返回值并不是t）
//执行defer方法，将t + 3 = 4
//函数返回返回值1
func DeferFunc2(i int) int { // 1
	t := i
	fmt.Printf("t1 %p value:%v \n", &t, t)
	defer func() {
		t += 3
		fmt.Printf("defer t2 %p value:%v \n", &t, t)
	}()
	return t
}
func DeferFunc22(i int) (result int) { // 1
	t := i
	fmt.Printf("t1 %p value:%v \n", &t, t)
	defer func() {
		t += 3
		fmt.Printf("defer t2 %p value:%v \n", &t, t)
	}()
	return t // 赋值 t=result  -> defer -> 返回
}

func DeferFunc3(i int) (t int) { // 3
	defer func() {
		t += i
	}()
	return 2
}

func DeferFunc4() (t int) { // 0, 2
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t) // 此时的t=0
	t = 1
	return 2
}

func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	DeferFunc4()
}
