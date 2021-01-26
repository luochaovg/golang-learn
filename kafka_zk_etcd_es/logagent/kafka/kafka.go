package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func Init(addr []string) (err error) {
	config := sarama.NewConfig()
	// tailf 包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel 返回

	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Print("producer closed , err:", err)
		return
	}

	return
}

// 往kafka发送数据 -> 生产者
func SendMsgToKafka(topic, data string) {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder(data)

	// send msg
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg faild , err:", err)
		return
	}

	fmt.Printf("partition:%v offset:%v \n ", partition, offset)
}
