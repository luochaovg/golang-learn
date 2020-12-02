package base

import "fmt"

func main() {

	age := 64

	if age > 18 {
		fmt.Println("你成年了")
	} else {
		fmt.Println("小屁孩")
	}

	if age := 23; age > 2 { // 注意： age 变量只在if里面生效
		fmt.Println("哈哈哈")
	}

	//for i := 0; i < 3; i++ {
	//	fmt.Println(i)
	//}

	i := 5
	for ; i < 10; i++ {
		fmt.Println(i)
		if i == 8 {
			//break
			//continue
		}
	}

	// 死循环
	//for {
	//	fmt.Println(1)
	//}

	s := "fasdfasdfafadsdfasf"
	for k, v := range s {
		fmt.Printf("%d %c \n", k, v)
	}

	n := 1
	switch n {
	case 2:
		fmt.Println("error")
		break
	case 3:
		fmt.Println("ok")
		break
	default:
		fmt.Println("error")
		break

	}

}
