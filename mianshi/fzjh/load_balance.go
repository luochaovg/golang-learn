package main

// 负载均衡 https://juejin.cn/post/6871169933150486542

// 负载均衡 - 随机负载

import (
	"errors"
	"math/rand"
	"time"
)

type RandomBalance struct {
	curIndex int
	rss      []string
}

func (r *RandomBalance) Add(client string) error {

	if len(client) == 0 {
		return errors.New("client is null")
	}
	r.rss = append(r.rss, client)

	return nil
}

func (r *RandomBalance) Get() (client string, err error) {

	if len(r.rss) == 0 {
		return "", errors.New("没有客户端连接")
	}

	rand.Seed(time.Now().UnixNano())
	curIndex := rand.Intn(len(r.rss))
	r.curIndex = curIndex

	client = r.rss[curIndex]

	return client, nil
}

func main() {

}
