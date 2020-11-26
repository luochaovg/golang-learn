package main

import "fmt"

type student struct {
	id   int64
	name string
}

// 造一个学生管理者
type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudents() {
	for _, v := range s.allStudent {
		fmt.Printf("学号:%d, 姓名:%s \n", v.id, v.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	var (
		id   int64
		name string
	)

	// 获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)

	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)

	// 根据用户输入创建结构体对象
	newStu := student{
		id:   id,
		name: name,
	}

	// 把新的学生放到s.allStudent这个map中
	s.allStudent[newStu.id] = newStu

	fmt.Println("添加成功！")
}

// 修改学生
func (s studentMgr) editStudeng() {

	var id int64
	var newName string

	// 获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)

	// 展示该学号对应的学生信息，如果没有提示查无此人
	stuObj, ok := s.allStudent[id]
	if !ok {
		fmt.Println("查无此学生！")
		return
	}

	fmt.Println("你要修改的学生名字：", stuObj.name)

	// 输入修改后的学生名
	fmt.Print("请输入学生的新名字：")
	fmt.Scanln(&newName)

	// 更新学生姓名
	stuObj.name = newName
	s.allStudent[id] = stuObj // 更新map中的学生

	fmt.Println("修改成功")
}

// 删除学生
func (s studentMgr) deleteStudent() {
	var id int64

	// 获取用户输入
	fmt.Print("请输入删除学生的学号：")
	fmt.Scanln(&id)

	// 如果没有提示查无此人
	_, ok := s.allStudent[id]
	if !ok {
		fmt.Println("查无此学生！")
		return
	}

	delete(s.allStudent, id)
	fmt.Println("删除成功")
}
