package main

import (
	"bufio"
	"fmt"
	"io"
	proto "liwenzhou/protocol"
	"net"
)

func Process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)

		bytes, _ := proto.Encode(msg)
		conn.Write(bytes) // 发送数据
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30001")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go Process(conn)
	}
}
