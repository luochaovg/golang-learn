package main

import (
	"fmt"
	"strconv"
	"strings"
)

//var (
//	rdb2 *redis.Client
//)
//var ctx2 = context.TODO()

//func init() {
//	rdb2 = redis.NewClient(&redis.Options{
//		Addr:     "172.16.13.143:6379",
//		Password: "lc910112",
//		DB:       0,
//		PoolSize: 100,
//	})
//
//	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancelFunc()
//
//	_, err := rdb2.Ping(timeout).Result()
//	if err != nil {
//		panic(err)
//	}
//}

func main() {

	//b, err := SaveStringToRedis("tst", "luochao", time.Hour)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(b)

	//value, err :=  rdb2.MGet(ctx2, []string{"name", "age", "abc"}...).Result()
	//if err != nil {
	//
	//}
	//fmt.Println(value)
	//
	//pipe := rdb2.Pipeline()
	//pipe.SetEX()
	//pipe.Exec(ctx2)

	sql := `
	select a, b, c from 
     luochao你好
    where a in (? , ?) 
	and b = ?
`

	fmt.Println(len(sql))
	fmt.Println(strings.Index(sql, "?"))

	tmpSql := ""
	index := 0
	for i, v := range sql {
		fmt.Println(i, string(v))
		if string(v) == "?" {
			tmpSql = tmpSql + "$" + strconv.Itoa(index+1)
			index++
		} else {
			tmpSql += string(v)
		}
	}

	fmt.Println("TMPSQL", tmpSql)
	//ns := strings.Replace(sql,"?","$",-1)
	//
	//fmt.Println(ns)
	return

	//api := &CacheMetricsByDateList{
	//	GameCode:  "test", // 游戏game_code
	//	Filter:    []string{"aa"}, // 过滤条件
	//	CacheUsed: true, // 是否走缓存，
	//	Metircs:   Revenue, // 需要获取的指标
	//	Expire:    60*time.Second, // 缓存TTL
	//	Date: "2025",
	//}
	//
	//val, err:= api.MerticsWapper(api.RevenueMetricLogic).RevenewUnseris()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(val)
}
