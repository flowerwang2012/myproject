package main

import (
	"github.com/coreos/etcd/clientv3"
	"os"
	"os/signal"
	"fmt"
	"time"
	"context"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	/*** code begin ***/

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379","127.0.0.1:2389","127.0.0.1:2399"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	kv := clientv3.NewKV(cli)
	putResp, err := kv.Put(context.TODO(), "/test/a", "sam")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("putResp: %+v\n", putResp)
	// 再写一个孩子
	kv.Put(context.TODO(),"/test/b", "another")
	// 再写一个同前缀的干扰项
	kv.Put(context.TODO(), "/testxxx", "干扰")

	getResp, err := kv.Get(context.TODO(), "/test/a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(getResp.Kvs)

	rangeResp, err := kv.Get(context.TODO(), "/test/", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rangeResp.Kvs)
	/*** code end ***/

	select {
	case <-ch:
		fmt.Println("退出程序")
		os.Exit(2)
	}
}
