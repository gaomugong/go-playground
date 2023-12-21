package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/gaohongsong/go-playground/go-with-test/consul/service"
	"github.com/gaohongsong/go-playground/go-with-test/etcd/regdiscover"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "pitou", "username parameter")

func main() {

	flag.Parse()

	// consulClient, err := capi.NewClient(capi.DefaultConfig())
	// if err != nil {
	// 	log.Fatalf("create consul client error: %s", err)
	// }
	//
	// entries, meta, err := consulClient.Health().Service("demo-server", "grpc", true, nil)
	// log.Printf("meta=%v, entries=%v, err=%v", meta, len(entries), err)
	// if err != nil {
	// 	log.Fatalf("service not found: %s", err)
	// }
	//
	// log.Println("Health endpoints:")
	// for _, entry := range entries {
	// 	fmt.Printf("%s:%d\n", entry.Service.Address, entry.Service.Port)
	// }

	// selectIndex := rand.Intn(len(entries))
	// entry := entries[selectIndex]
	// svrUrl := fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port)

	svrDiscover := regdiscover.NewServiceDiscovery([]string{"127.0.0.1:2379"})
	defer svrDiscover.Close()

	_ = svrDiscover.WatchService("/web")

	for {
		select {
		case <-time.Tick(time.Second * 3):
			entries := svrDiscover.GetServices()
			log.Println("discover services:", entries)

			selectIndex := rand.Intn(len(entries))
			svrUrl := entries[selectIndex]
			log.Printf("target-grpc -> %s\n", svrUrl)

			conn, err := grpc.Dial(svrUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("connect server error: %s", err)
			}
			defer conn.Close()

			client := service.NewDemoClient(conn)

			resp, err := client.SendRequest(context.Background(), &service.Request{Username: *name})
			if err != nil {
				log.Fatalf("send request error: %s", err)
			}
			log.Printf("response message: %v", resp.GetMessage())
		}
	}
}
