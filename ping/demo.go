package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

func main() {
	rst := tcpPort("43.231.196.71", "3389")
	fmt.Println("tcp port", rst)

	rst2 := ping2("43.231.196.71")
	fmt.Println("ping", rst2)
}

func ping2(ip string) bool {
	// -c 发包数量
	cmd := exec.Command("ping", ip, "-c", "4", "-W", "5")
	//fmt.Println("NetWorkStatus Start:", time.Now().Unix())
	err := cmd.Run()
	//fmt.Println("NetWorkStatus End  :", time.Now().Unix())
	if err != nil {
		return false
	}

	return true
}

func tcpGather(ip string, ports []string) map[string]string {
	// 检查 emqx 1883, 8083, 8080, 18083 端口

	results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results[port] = "failed"
		} else {
			if conn != nil {
				results[port] = "success"
				_ = conn.Close()
			} else {
				results[port] = "failed"
			}
		}
	}
	return results
}

func tcpPort(ip string, port string) bool {
	address := net.JoinHostPort(ip, port)

	// 3 秒超时
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)

	if err != nil {
		return false
	}

	if conn != nil {
		_ = conn.Close()
		return true
	} else {
		return false
	}

}
