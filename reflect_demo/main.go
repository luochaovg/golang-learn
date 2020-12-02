package main

import (
	"fmt"
	"reflect"
)

type myInt int16

type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type Kind: %v\n", v.Name(), v.Kind())

}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// 反射
// 接口类型的变量底层是分为两部分： 动态类型和动态值
// 反射的应用 json等数据解析/ORM等工具

func main() {
	var a float32 = 3.14
	//reflectType(a) // type:float32
	reflectValue(a)
	var b int64 = 100
	reflectValue(b)
	//reflectType(b) // type:int64
	//
	//var c myInt = 23
	//reflectType(c)
	//
	//var d = Cat{}
	//reflectType(d)
	//reflectSetValue1(b)
	reflectSetValue2(&b)
	fmt.Println(b)

}
