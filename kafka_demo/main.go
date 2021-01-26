package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()
	// tailf 包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel 返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"}, config)
	if err != nil {
		fmt.Print("producer closed , err:", err)
		return
	}

	defer client.Close()

	// send msg
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg faild , err:", err)
		return
	}

	fmt.Printf("pid:%v offset:%v \n ", pid, offset)
}
