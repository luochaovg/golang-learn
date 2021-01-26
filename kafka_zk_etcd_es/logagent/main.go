package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"liwenzhou/kafka_zk_etcd_es/logagent/config"
	"liwenzhou/kafka_zk_etcd_es/logagent/kafka"
	"liwenzhou/kafka_zk_etcd_es/logagent/taillog"
	"time"
)

var (
	cfg = new(config.AppConf)
)

// logAgnet 入口
func main() {
	// 0.加载配置文件
	//cfg, err := ini.Load("./config/config.ini")

	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Println("ini init failed , err:", err)
		return
	}

	address := cfg.Address
	logPath := cfg.FileName
	//fmt.Println(address, logPath)
	//return

	// 1.初始化kafka链接
	err = kafka.Init([]string{address})
	if err != nil {
		fmt.Println("kafka init failed , err:", err)
		return
	}
	fmt.Println("init kafaka success")

	// 2.打开日志文件准备收集日志
	err = taillog.Init(logPath)
	if err != nil {
		fmt.Println("taillog init failed , err:", err)
		return
	}

	fmt.Println("init taillog success")

	// 执行
	//taillog.ReadLog()
	run()
}

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SendMsgToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
