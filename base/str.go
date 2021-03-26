package main

import (
	"fmt"
	"unsafe"
)

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}
func main() {

}

func main_demo() {
	str := "hello"
	println([]byte(str))
}
