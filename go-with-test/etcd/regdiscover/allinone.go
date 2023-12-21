package main

import (
	"context"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var endpoints = []string{
	"127.0.0.1:2379",
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// =============== 服务注册部分 ===================
func RegisterServiceAndWatch(key, value string, leaseTtl int64) (*clientv3.Client, <-chan *clientv3.LeaseKeepAliveResponse, clientv3.LeaseID) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}

	// 创建租约
	grant, err := cli.Grant(context.Background(), leaseTtl)
	if err != nil {
		panic(err)
	}

	// 注册服务并绑定租约
	if _, err = cli.Put(context.Background(), key, value, clientv3.WithLease(grant.ID)); err != nil {
		panic(err)
	}

	// 设置租约时间
	aliveRespChan, err := cli.KeepAlive(context.Background(), grant.ID)
	if err != nil {
		panic(err)
	}

	return cli, aliveRespChan, grant.ID
}

func WatchAliveChan(aliveRespChan <-chan *clientv3.LeaseKeepAliveResponse) {
	for leaseKeepResp := range aliveRespChan {
		log.Println("续约成功", leaseKeepResp)
	}
	log.Printf("关闭续约")
}

// =============== 服务发现部分 ===================
var (
	serviceList = make(map[string]string)
	lock        sync.Mutex
)

func PutSetService(kv *mvccpb.KeyValue) {
	lock.Lock()
	defer lock.Unlock()
	key, value := string(kv.Key), string(kv.Value)
	serviceList[key] = value
	log.Printf("put key=%s, v=%s", key, value)
}

func DeleteService(kv *mvccpb.KeyValue) {
	lock.Lock()
	defer lock.Unlock()
	key, value := string(kv.Key), string(kv.Value)
	delete(serviceList, key)
	log.Printf("delete key=%s, v=%s", key, value)
}

func GetAndPrintServices() {
	lock.Lock()
	defer lock.Unlock()
	log.Printf("Services: %v", serviceList)
}

func DiscoverServiceAndWatch(key string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	res, err := cli.Get(context.Background(), key, clientv3.WithPrefix())
	panicErr(err)

	for _, ev := range res.Kvs {
		serviceList[string(ev.Key)] = string(ev.Value)
	}

	// 监听服务的变化
	watchChan := cli.Watch(context.Background(), key, clientv3.WithPrefix())
	for watchResp := range watchChan {
		for _, ev := range watchResp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				PutSetService(ev.Kv)
			case mvccpb.DELETE:
				DeleteService(ev.Kv)
			}
		}
	}
}

func main() {
	// 注册服务
	cli, watchChan, leaseId := RegisterServiceAndWatch("/web", "127.0.0.1:8000", 5)

	// 注销服务
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		log.Println("撤销租约")
		_, _ = cli.Revoke(context.Background(), leaseId)
		_ = cli.Close()
		cancel()
	}()

	go WatchAliveChan(watchChan)

	// 模拟服务发现
	go DiscoverServiceAndWatch("/web")

	go func() {
		// 模拟服务发现
		for {
			select {
			case <-time.Tick(time.Second * 5):
				GetAndPrintServices()
			case <-ctx.Done():
				log.Println("模拟服务发现退出")
			}
		}
	}()

	select {
	case t := <-time.After(time.Second * 10):
		log.Printf("server exited, t=%v", t)
	}

}
