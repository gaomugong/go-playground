package main

import (
	"context"
	"flag"
	"fmt"
	"log"

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

	heath, info, err := consulClient.Agent().AgentHealthServiceByID("demo-server")
	log.Printf("heath=%v, info.Service=%v, err=%v", heath, info.Service, err)
	if err != nil {
		log.Fatalf("service not found: %s", err)
	}

	svrUrl := fmt.Sprintf("%s:%d", info.Service.Address, info.Service.Port)
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
