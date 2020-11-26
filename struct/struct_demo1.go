package _struct

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看/新增/删除学生
*/

type student struct {
	id   int64
	name string
}

var allStudent map[int64]*student // 变量声明

func main() {
	allStudent = make(map[int64]*student, 48) // 初始化，开辟内存空间

	for {
		// 1.打印菜单
		fmt.Println("welecome to student manger systerm")
		fmt.Println(`
		     1,查看所有学生
			 2,新增学生
		     3,删除学生
			 4,退出`)
		fmt.Println("请输入你要干啥：")

		// 2.等待用户选择下一步
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)

		// 3.执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
			break
		case 2:
			addStudent()
			break
		case 3:
			deleteStudent()
			break
		case 4:
			os.Exit(1)
		default:
			fmt.Println("gun")
		}
	}
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	for _, v := range allStudent {
		fmt.Printf("你的学号%d, 姓名：%s \n", v.id, v.name)
	}
}
func addStudent() {
	// 1.创建一个学生
	var (
		id   int64
		name string
	)

	fmt.Print("请输入学生的学号")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名")
	fmt.Scanln(&name)

	// 2.追加到allStudent这个map
	allStudent[id] = newStudent(id, name)
}

func deleteStudent() {
	var id int64
	fmt.Print("请输入要删除的学生学号")
	fmt.Scanln(&id)
	delete(allStudent, id)
}
