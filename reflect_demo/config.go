package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0.参数的校验
	// 0.1 传入的data参数必须是指针类型
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		//err = fmt.Errorf("data should be a pointer") // 格式化输出之后返回一个error类型
		err = errors.New("data should be a pointer") // 创建一个错误
		return
	}

	// 0.2 传进来的data参数必须是结构体类型指针，因为配置文件中各种健值对赋值给结构体字段
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct pointer") // 创建一个错误
		return
	}

	// 1.读文件得到字节类型的数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	lineSlice := strings.Split(string(b), "\n") // 将字节类型的文件内容转换为字符串
	//fmt.Printf("%#v", lineSlice)

	// 2.一行一行的读数据
	var structName string

	for idx, line := range lineSlice {
		// 去掉首位空格
		line = strings.TrimSpace(line)

		// 空行跳过
		if len(line) == 0 {
			continue
		}

		// 2.1 如果是注释就忽略
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		// 2.2 如果是[开头就表示是节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			// 把这一行首尾[]去掉，拿到中间内容把首尾空格去掉
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			// 根据字符串sectionName 去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("%s , %s \n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是[开头就是=分隔的健值对
			// 2.3.1 等号分隔，左边key,右边值
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax err", idx+1)
				return
			}

			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			// 2.3.2 根据structName 去data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}

			// 2.3.3 遍历结构体的每一个字段，判断tag是不是等于key
			var fieldName string
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) // tag 信息是存储在类型信息中的
				if filed.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = filed.Name
				}
			}

			// 2.3.4 如果key=tag， 给这个字段赋值
			// 2.3.4.1 根据fieldName 去取出这个字段
			if len(fieldName) == 0 { // 在结构体找不到对应的字段
				continue
			}

			fileObj := sValue.FieldByName(fieldName)
			// 2.3.4.2 对其赋值
			fmt.Println(fieldName, fileObj.Type())

			switch fileObj.Type().Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}
	}

	return
}

func main() {
	var cfg Config

	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed , err:%v\n", err)
		return
	}

	//
	fmt.Printf("%#v\n", cfg)
}
