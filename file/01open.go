package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile() {
	fileObj, err := os.Open("./text.txt")

	if err != nil {
		fmt.Println("open file err", err)
		return
	}

	defer fileObj.Close()

	size := 4
	var tmp = make([]byte, size) // 指定读的长度
	for {
		n, err := fileObj.Read(tmp)

		if err == io.EOF {
			fmt.Println("读完了")
			return
		}

		if err != nil {
			fmt.Printf("read from file fail err %v \n", err)
		}

		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))

		if n < size {
			return
		}
	}
}

// 利用bufio 这个包读文件
func readFromFileByBufio() {
	fileObj, err := os.Open("./text.txt")

	if err != nil {
		fmt.Println("open file err", err)
		return
	}

	defer fileObj.Close()

	// TODO 因为 fileObj 实现了 Read()方法, 而io.Reader是一个接口， 接口只有一个Read()方法

	reader := bufio.NewReader(fileObj)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("没有了")
			return
		}

		if err != nil {
			fmt.Printf("Read Line 出错了, %v \n", err)
			return
		}

		fmt.Print(line)
	}
}

// 全部读
func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./text.txt")
	if err != nil {
		fmt.Printf("read file err %v\n", err)
	}
	fmt.Println(string(ret))
}

func main() {
	//readFromFileByBufio()
	readFromFileByIoutil()
}
