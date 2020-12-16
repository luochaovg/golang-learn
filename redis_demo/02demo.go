package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	rdb2 *redis.Client
)

func init() {
	rdb2 = redis.NewClient(&redis.Options{
		Addr:     "192.168.158.88:6379",
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := rdb2.Ping(timeout).Result()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx2 := context.Background()

	err := rdb2.Set(ctx2, "name", "luochao2222", 0).Err()
	if err != nil {
		panic(err)
	}
}
