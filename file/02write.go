package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件写内容
// 位运算符 "|" 或的意思
func openFileAndWrite() {
	fileObj, err := os.OpenFile("./write.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Printf("open file err %v \n", err)
		return
	}

	defer fileObj.Close()

	//line := []byte{'a', 'c'}
	//line := []byte("发顺丰寄啦开始放假 \n")
	//fmt.Printf("%T, %#v \n", line, line)

	//n, err := fileObj.Write(line)
	//if err != nil {
	//	fmt.Printf("write err %v\n", err)
	//}
	//fmt.Println(n)

	aa := "你发那是否健康阿水煎服啦 \n"
	n, err := fileObj.WriteString(aa)
	if err != nil {
		fmt.Printf("write err %v\n", err)
	}
	fmt.Println(n)

}

func writeBufio() {

	var fileObj *os.File
	var err error

	fileObj, err = os.OpenFile("./write.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Printf("open file err %v \n", err)
		return
	}

	defer fileObj.Close()

	wr := bufio.NewWriter(fileObj)
	n, err := wr.WriteString("abdfasdf \n") // 写到缓存
	wr.Flush()                              // 将缓存写到文件
	if err != nil {
		fmt.Printf("write err %v\n", err)
	}
	fmt.Println(n)

}

func main() {
	//openFileAndWrite()
	writeBufio()

	// fmt.Fprintln(os.Stdout, "aaaaaaaa")

}
