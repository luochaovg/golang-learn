package main

import (
	"sync"
)

// sync.Once  并发安全

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
