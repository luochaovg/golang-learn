package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
	"github.com/jmoiron/sqlx"          // 依赖 go-sql-driver/mysql
)

// sqlx demo
// https://www.liwenzhou.com/posts/Go/sqlx/

var db *sqlx.DB // 是一个连接池对象

// TODO  注意大写
type user struct {
	ID       int
	Username string
	Password string
}

func initDB() (err error) {
	dsn := "root:lc910112@tcp(192.168.158.88:3306)/lez-analysis"
	// 不会校验用户名密码正确
	db, err = sqlx.Connect("mysql", dsn) // 这里的db 就是全局db

	if err != nil { // dsn 格式不正确的时候会报错
		//fmt.Printf("dsn: %s invalid, err:%v\n", dsn, err)
		return
	}

	// 设置数据连接池的最大连接数,
	//  注意：如果业务里面没有释放链接，程序会卡住
	db.SetMaxOpenConns(10)

	// 最大空闲连接数
	db.SetMaxIdleConns(4)

	//fmt.Println(db)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed , err:%v\n", err)
		return
	}

	fmt.Println("连接数据库成功")

	// 查一条
	sqlStr := "select id,username,password from admin_users where id=?;"
	var u user
	db.Get(&u, sqlStr, 1) // TODO 注意
	fmt.Println(u.ID)

	// 查多条
	var userList []user
	sqlStr2 := "select id,username,password from admin_users where id>?;"
	err = db.Select(&userList, sqlStr2, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(userList))
	fmt.Printf("%#v\n", userList)
}
