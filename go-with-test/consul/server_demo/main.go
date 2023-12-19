package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/gaohongsong/go-playground/go-with-test/consul/service"
	"google.golang.org/grpc"

	capi "github.com/hashicorp/consul/api"
)

var port = flag.Int("port", 8000, "tcp bind port")

func main() {
	flag.Parse()

	// create grpc server
	svr := grpc.NewServer()

	// create tcp listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	// register demo server to grpc server
	// make sure to implement service first
	service.RegisterDemoServer(svr, &DemoServerImpl{})

	client, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		log.Fatalf("create consul client error: %s", err)
	}

	err = client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
		ID:      "demo-server",
		Name:    "demo-server",
		Tags:    []string{"grpc", "demo-server"},
		Address: "127.0.0.1",
		Port:    *port,
	})
	if err != nil {
		log.Fatalf("register demo service to consul error: %s", err)
	}

	log.Printf("start demo server at :%d", *port)
	if err := svr.Serve(lis); err != nil {
		log.Fatalf("start demo server failed: %s", err)
	}
}

// DemoServerImpl demo server interface implement
type DemoServerImpl struct {
	service.UnimplementedDemoServer
}

func (ds DemoServerImpl) SendRequest(ctx context.Context, req *service.Request) (*service.Response, error) {
	username := req.GetUsername()
	log.Printf("received username: %s\n", username)
	return &service.Response{Message: "hello, " + username}, nil
}
