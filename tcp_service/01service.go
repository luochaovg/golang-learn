package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp/server/main.go

// TCP server端

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	// 1.本地断开启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	// 2.建立连接
	//conn, err := listen.Accept()
	//if err != nil {
	//	fmt.Println("accept failed , err:", err)
	//	return
	//}

	// 3.与客户端链接
	//var tmp [128]byte
	//n, err := conn.Read(tmp[:])
	//if err != nil {
	//	fmt.Println("read from conn failed, err:", err)
	//	return
	//}
	//fmt.Println(string(tmp[:n]))

	for {
		// 2.建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 3.启动一个goroutine处理连接
		go process(conn)
	}

}
