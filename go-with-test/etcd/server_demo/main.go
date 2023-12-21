package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/gaohongsong/go-playground/go-with-test/consul/service"
	"github.com/gaohongsong/go-playground/go-with-test/etcd/regdiscover"
	"google.golang.org/grpc"
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

	// grpc_health_v1.RegisterHealthServer(svr, health.NewServer())

	// register demo server to grpc server
	// make sure to implement service first
	service.RegisterDemoServer(svr, &DemoServerImpl{})

	// config := capi.DefaultConfig()
	// config.Address = consulAddress
	// client, err := capi.NewClient(config)
	// if err != nil {
	// 	log.Fatalf("create consul client error: %s", err)
	// }
	//
	// err = client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
	// 	Name:    "demo-server",
	// 	ID:      "demo-server-" + strconv.Itoa(*port),
	// 	Tags:    []string{"grpc", "demo-server"},
	// 	Address: "127.0.0.1",
	// 	Port:    *port,
	// 	Check: &capi.AgentServiceCheck{
	// 		// TCP:  "127.0.0.1:" + strconv.Itoa(*port),
	// 		GRPC:                           "127.0.0.1:" + strconv.Itoa(*port), // 健康检查地址只需要写grpc服务地址端口
	// 		Interval:                       "5s",
	// 		Timeout:                        "5s",
	// 		DeregisterCriticalServiceAfter: "30s", // 故障检查失败30s后 consul自动将注册服务删除
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("register demo service to consul error: %s", err)
	// }

	portStr := strconv.Itoa(*port)

	// 创建基于etcd的服务注册中心
	svrReg, err := regdiscover.NewServiceRegister([]string{"127.0.0.1:2379"}, 5)
	if err != nil {
		log.Fatalln(err)
	}

	// 注册服务
	if err = svrReg.RegisterService("/web/"+portStr, "127.0.0.1:"+portStr); err != nil {
		log.Fatal(err)
	}

	// 服务续约
	go svrReg.KeepServiceAlive()

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

func (ds DemoServerImpl) Ping(ctx context.Context, req *service.PingRequest) (*service.Response, error) {
	return &service.Response{Message: "pong"}, nil
}
