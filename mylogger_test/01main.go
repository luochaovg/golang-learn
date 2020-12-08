package main

import (
	"liwenzhou/mylogger"
	"time"
)

var log mylogger.Logger // 声明一个全局的接口变量

// 测试我们自己的写的日志库
func main() {

	log = mylogger.NewConsoleLog("debug") // 终端日志实列
	//log = mylogger.NewFileLogger("Info", "./", "law.log", 10*1024*1024) // 文件日志实列

	for {
		id := 1000
		name := "law"
		log.Debug("这是一条debug日志 id:%d, name:%s", id, name)
		log.Trace("这是一条trace日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条waring日志")
		log.Error("这是一条error日志")
		log.Fatal("这是一条fatal日志")

		time.Sleep(time.Second)
	}
}
