package main

import "fmt"

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func loadIni(v interface{}) {

}

func main() {
	var mc MysqlConfig
	loadIni(&mc)

	fmt.Println(mc, mc.Address)
}
