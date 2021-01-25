package main

import (
	"fmt"
	"sync"
)

// 并发安全的单例模式

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

// 懒汉模式 : 是开源项目中使用最多的一种，最大的缺点是非线程安全的

type singleton2 struct{}

var instance2 *singleton2

func GetInstance2() *singleton2 {
	if instance2 == nil {
		instance2 = &singleton2{}
	}

	return instance2
}
