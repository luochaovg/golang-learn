package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student3 struct {
}

func (s *Student3) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	//var peo People
	//var stu = Student3{}
	//peo = stu
	var peo People = Student3{}

	think := "bitch"
	fmt.Println(peo.Speak(think))

}
