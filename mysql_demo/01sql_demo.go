package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// 下载驱动
//  go get github.com/go-sql-driver/mysql
// go get 包的路径就是下载第三方依赖
// 将第三方的依赖默认保存在	$GOPATH/src

var db *sql.DB // 是一个连接池对象

type user struct {
	id       int
	username string
	password string
}

func initDB() (err error) {
	dsn := "root:lc910112@tcp(192.168.158.88:3306)/lez-analysis"
	// 不会校验用户名密码正确
	db, err = sql.Open("mysql", dsn) // 这里的db 就是全局db
	if err != nil {                  // dsn 格式不正确的时候会报错
		//fmt.Printf("dsn: %s invalid, err:%v\n", dsn, err)
		return
	}

	// 尝试与数据库建立连接
	err = db.Ping()
	if err != nil {
		//fmt.Printf("open %s failed, err:%v\n", dsn, err)
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

// 单条查询
func queryRowDemo(id int) user {
	// 1. sql 语句
	sqlStr := "select id,username,password from admin_users where id=?;"
	var u user

	// 2.执行 , 从连接池里哪一个链接出来去数据库查询单条记录
	rowObj := db.QueryRow(sqlStr, id)

	// 3. 拿到结果
	// TODO 注意：必须对rowObj 对象调用Scan方法， 因为该方法会释放数据库链接
	err := rowObj.Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Println(err)
	}

	// 2/3合并， 推荐
	// db.QueryRow(sqlStr, 2).Scan(&u.id, &u.username, &u.password)

	fmt.Printf("u:%#v\n", u)
	return u
}

// 多行查询
func queryMultiRowDemo(n int) {
	// 1. sql 语句
	sqlStr := "select id,username,password from admin_users where id>?;"

	// 2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}

	//Todo 3.非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 4.循环取值
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}

		fmt.Printf("u:%#v\n", u)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed , err:%v\n", err)
	}

	fmt.Println("连接数据库成功")

	//var u1 user
	//u1 = queryRowDemo(2)
	//fmt.Println(u1)

	queryMultiRowDemo(0)

}

// TODO Mysql 预处理
// 预处理查询示例
func prepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id)
	}
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	//var m = map[string]int{
	//	"张三": 12,
	//	"李四": 14,
	//}
	//for k, v := range m {
	//	_, err = stmt.Exec(k, v)
	//}

	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

// TODO 事务
// 事务操作示例
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}
