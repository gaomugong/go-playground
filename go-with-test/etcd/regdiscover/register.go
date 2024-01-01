package regdiscover

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// ServiceRegister 创建租约注册服务
type ServiceRegister struct {
	cli     *clientv3.Client // etcd client
	leaseID clientv3.LeaseID // 租约ID
	// 租约keepalieve相应chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
}

// NewServiceRegister 新建注册服务
func NewServiceRegister(endpoints []string, ttl int64) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 申请租约设置时间keepalive并注册服务
	grant, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		log.Fatal(err)
	}

	// 设置续租 定期发送需求请求
	keepAliveChan, err := cli.KeepAlive(context.Background(), grant.ID)
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceRegister{
		cli:           cli,
		leaseID:       grant.ID,
		keepAliveChan: keepAliveChan,
	}, nil
}

// RegisterService 注册服务并绑定租约
func (s *ServiceRegister) RegisterService(key, value string) error {
	if _, err := s.cli.Put(context.Background(), key, value, clientv3.WithLease(s.leaseID)); err != nil {
		return err
	}
	return nil
}

// keepAlive 监听续租情况
func (s *ServiceRegister) KeepServiceAlive() {
	for aliveKeepResp := range s.keepAliveChan {
		log.Println("续约成功", aliveKeepResp)
	}
	log.Println("停止续租")
}

// Close 注销服务
func (s *ServiceRegister) Close() error {
	// 撤销租约
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	log.Println("撤销租约")
	return s.cli.Close()
}

func main1() {
	// 创建服务注册中心
	svrReg, err := NewServiceRegister([]string{"127.0.0.1:2379"}, 5)
	if err != nil {
		log.Fatalln(err)
	}

	// 注册服务
	if err = svrReg.RegisterService("/web", "127.0.0.1:8000"); err != nil {
		log.Fatal(err)
	}

	// 服务续约
	go svrReg.KeepServiceAlive()

	select {
	case <-time.After(20 * time.Second):
		svrReg.Close()
	}
}
