package main

import (
	"context"
	"flag"
	"log"
	"time"

	p "github.com/gaohongsong/go-playground/go-with-test/grpc/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name string
var addr string

func init() {
	log.Println("init")
	flag.StringVar(&name, "name", "", "name to greet")
	flag.StringVar(&addr, "addr", "localhost:50051", "grpc server address")
}

func main() {
	flag.Parse()
	log.Printf("\nname: %s\naddr: %s\n", name, addr)

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	// Invoke rpc method
	defer conn.Close()
	c := p.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &p.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("greet failed: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
