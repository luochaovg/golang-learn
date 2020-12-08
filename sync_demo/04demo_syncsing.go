package main

import "sync"

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
