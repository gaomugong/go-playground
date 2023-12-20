package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"

	"github.com/gaohongsong/go-playground/go-with-test/consul/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	capi "github.com/hashicorp/consul/api"
)

var name = flag.String("name", "pitou", "username parameter")

func main() {

	flag.Parse()

	consulClient, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		log.Fatalf("create consul client error: %s", err)
	}

	entries, meta, err := consulClient.Health().Service("demo-server", "grpc", true, nil)
	log.Printf("meta=%v, entries=%v, err=%v", meta, len(entries), err)
	if err != nil {
		log.Fatalf("service not found: %s", err)
	}

	log.Println("Health endpoints:")
	for _, entry := range entries {
		fmt.Printf("%s:%d\n", entry.Service.Address, entry.Service.Port)
	}

	selectIndex := rand.Intn(len(entries))
	entry := entries[selectIndex]
	svrUrl := fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port)
	log.Printf("target-grpc -> %s\n", svrUrl)

	conn, err := grpc.Dial(svrUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect server error: %s", err)
	}
	defer conn.Close()

	client := service.NewDemoClient(conn)
	// pong, err := client.Ping(context.Background(), &service.PingRequest{})
	// if err != nil {
	// 	log.Fatalf("send ping request error: %s", err)
	// }
	// log.Printf("ping -> %s", pong.GetMessage())

	resp, err := client.SendRequest(context.Background(), &service.Request{Username: *name})
	if err != nil {
		log.Fatalf("send request error: %s", err)
	}
	log.Printf("response message: %v", resp.GetMessage())
}
