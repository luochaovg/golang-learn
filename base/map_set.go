package main

import "fmt"

// 集合的特性 无序/key唯一

type Collect map[string]struct{}

func (c Collect) Has(key string) bool {
	_, ok := c[key]
	return ok
}

func (c Collect) Add(key string) {
	c[key] = struct{}{}
}

func (c Collect) Delete(key string) {
	delete(c, key)
}

func main() {
	set := make(Collect)

	set.Add("LC")
	set.Add("XX")

	fmt.Printf("set is : %#v", set)

	fmt.Println(set.Has("LC"))
	fmt.Println(set.Has("LX"))
}
