package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect etcd faild , err:%v \n", err)
		return
	}
	defer cli.Close()

	// put
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//_, err = cli.Put(ctx, "name", "luochao")
	//cancel()
	//if err != nil {
	//	fmt.Printf("put to etcd failed, err:%v\n", err)
	//	return
	//}

	// get
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, "name")
	//cancel()
	//if err != nil {
	//	fmt.Printf("get from etcd failed , err:%v\n", err)
	//	return
	//}
	//
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s:%s \n", ev.Key, ev.Value)
	//}

	// watch
	rch := cli.Watch(context.Background(), "name") // <- chan WatchChan
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%s key:%s value:%s \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
