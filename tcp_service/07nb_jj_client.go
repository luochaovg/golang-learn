package main

import (
	"bufio"
	"fmt"
	proto "liwenzhou/protocol"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30001")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	//defer conn.Close()
	//for i := 0; i < 20; i++ {
	//	msg := `Hello, Hello. How are you?`
	//	data, err := proto.Encode(msg)
	//	if err != nil {
	//		fmt.Println("encode msg failed, err:", err)
	//		return
	//	}
	//	conn.Write(data)
	//}

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("请说话：")
		//if msg2 != "" {
		//	fmt.Print("服务端回复：" + msg2)
		//}

		//fmt.Scanln(&msg)

		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}

		if msg == "exit" {
			break
		}

		conn.Write(data)

		reader1 := bufio.NewReader(conn)
		msg2, err := proto.Decode(reader1)

		//从服务端接收回复的消息
		//var buf [1024]byte
		//n, err := conn.Read(buf[:])
		//if err != nil {
		//	fmt.Printf("read failed,err:%v\n", err)
		//	return
		//}

		fmt.Printf("收到服务端回复:%v\n", msg2)
	}

	conn.Close()
}
