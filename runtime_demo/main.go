package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	getInfo()
}

func getInfo() {

	// 运行时
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() failed , err:\n ")
	}

	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file) // runtime/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line) // 11
}
