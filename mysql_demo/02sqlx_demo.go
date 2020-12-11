package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed , err:%v\n", err)
	}

	fmt.Println("连接数据库成功")

	//var u1 user
	//u1 = queryRowDemo(2)
	//fmt.Println(u1)

}
