package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp cliecnt

func main() {
	//b := []byte("hello world")
	//fmt.Println(b)
	//fmt.Printf("%#v, %T \n", b, b)
	//	1.建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("daial 127.0.0.1:20000 failed , err", err)
	}

	// 2.发送数据
	//var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请说话：")
		//fmt.Scanln(&msg)

		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}

		conn.Write([]byte(msg))
	}

	conn.Close()
}
