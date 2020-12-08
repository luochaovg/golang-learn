package main

import (
	"fmt"
	proto "liwenzhou/protocol"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30001")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
