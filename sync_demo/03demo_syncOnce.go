package main

import (
	"sync"
)

// sync.Once  并发安全

// 使用场景
// 某些函数值需要执行一次的时候，就可以使用sync.Once

var icons map[string]int
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]int{
		"left":  1,
		"up":    2,
		"right": 3,
		"down":  4,
	}
}

// Icon 是并发安全的
func Icon(name string) int {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main() {

}
