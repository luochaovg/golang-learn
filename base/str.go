package main

import (
	"fmt"
)

const cl = 100

var bl = 200

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	temp := "http://"
	fmt.Println([]byte(temp))
}

func main_demo() {
	str := "hello"
	println([]byte(str))

}
