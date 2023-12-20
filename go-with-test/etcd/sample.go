package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	// 创建客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})

	if err != nil {
		log.Fatal("new client error: ", err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	key, value := "hello", "world"
	putResp, err := cli.Put(ctx, key, value)
	cancel()

	if err != nil {
		log.Printf("error put: key=%s, err=%s", key, err)
	}
	log.Printf("put key=%s, value=%s, reps=%v", key, value, putResp)

	resp, err := cli.Get(context.TODO(), key)
	if err != nil {
		log.Printf("error get: key=%s, err=%s", key, err)
	}

	for _, kv := range resp.Kvs {
		log.Printf("get key=%s, value=%s, reps=%v", kv.Key, kv.Value, resp)
	}

	watchChan := cli.Watch(context.Background(), "hello")
	for resp := range watchChan {
		for _, e := range resp.Events {
			log.Printf("Type: %s, key: %s, value: %s", e.Type, e.Kv.Key, e.Kv.Value)
		}
	}
}
