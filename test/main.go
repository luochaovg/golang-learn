package main

import (
	"fmt"
	"liwenzhou/split_string"
)

func main() {
	ret := split_string.Split("babfdabfas", "b")
	fmt.Printf("%#v,\n", ret)
}
