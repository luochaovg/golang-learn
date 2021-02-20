package main

// Go语言中的单例模式
// https://www.liwenzhou.com/posts/Go/singleton_in_go/

import (
	"fmt"
	"sync"
)

type singleton struct {
	Test string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			Test: "hello world",
		}
	})

	return instance
}

func main() {
	t := GetInstance()
	fmt.Printf("%#v \n", t)
}
