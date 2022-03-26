package main

type S struct{}

var i interface{}

func f(x interface{}) {}

func g(x *interface{}) {}

func main() {
	s := S{}
	p := &s

	// TODO 看到这道题需要第一时间想到的是Golang是强类型语言，interface是所有golang类型的父类
	//  函数中func f(x interface{})的interface{}可以支持传入golang的任何类型，包括指针，
	//  但是函数func g(x *interface{})只能接受*interface{}

	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D

}
