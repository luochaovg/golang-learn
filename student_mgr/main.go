package main

import (
	"fmt"
	"os"
)

/**
学生管理系统
*/

var smr studentMgr

func showMenu() {
	fmt.Println("welecome sms")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 编辑学生
	4. 删除学生
	5. 退出
`)
}

func main() {

	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}

	for {
		showMenu()

		// 等待用户输入
		fmt.Print("请输入要执行的序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)

		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudeng()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("Gun")

		}
	}
}
