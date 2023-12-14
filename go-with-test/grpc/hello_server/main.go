package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	p "github.com/gaohongsong/go-playground/go-with-test/grpc/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	p.UnimplementedGreeterServer
}

func (server) SayHello(ctx context.Context, req *p.HelloRequest) (*p.HelloReply, error) {
	name := req.GetName()
	log.Printf("Received: %s", name)
	return &p.HelloReply{Message: fmt.Sprintf("hello %s", name)}, nil
}

// 这里返回的是个指针，别被坑到了...
var port = flag.Int("port", 50051, "The server port")

func main() {
	flag.Parse()

	log.Printf("%d", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gServer := grpc.NewServer()
	p.RegisterGreeterServer(gServer, &server{})
	log.Printf("server listen at %v", lis.Addr())

	if err := gServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
