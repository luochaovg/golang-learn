package main

import (
	"fmt"
	"runtime"

	//"runtime"
	"sync"
)

type student struct {
	Name string
	Age  int
}

//func main() {
//}

func pase_student() {
	m := make(map[string]*student)

	stus := []student{
		{
			Name: "lc",
			Age:  1,
		},
		{
			Name: "xx",
			Age:  2,
		},
	}

	// TODO golang 的  for ... range 语法中， stu 变量会被复⽤，每次循环会将集合中的值复制给这个变量，因此，会导致最后 m 中的 map 中储存的都是 stus 最后⼀个 student
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	fmt.Printf("m : %#v", m)
}

func main() {
	//这个输出结果决定来⾃于调度器优先调度哪个G。从runtime的源码可以看到，当创建⼀
	//个G时，会优先放⼊到下⼀个调度的 runnext 字段上作为下⼀次优先调度的G。因此，
	//最先输出的是最后创建的G，也就是9.
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("ii: ", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
