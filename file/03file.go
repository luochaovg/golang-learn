package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./text.txt")

	if err != nil {
		fmt.Printf("open file err %v", err)
		return
	}
	// 1  文件对象的类型
	fmt.Printf("%T\n", fileObj)

	// 2 获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file err %v", err)
		return
	}
	fmt.Printf("file size : %dB \n", fileInfo.Size())

}
