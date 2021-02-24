package main

// Go语言中的单例模式 ： 单列： 只实例化一次
// https://www.liwenzhou.com/posts/Go/singleton_in_go/

// php版本：
// https://phpenthusiast.com/blog/the-singleton-design-pattern-in-php
// https://juejin.cn/post/6844903657385754638

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

	// 如果已经实例化就返回
	if instance.Test != "" {
		return instance
	}

	// 并发安全 实例化
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
