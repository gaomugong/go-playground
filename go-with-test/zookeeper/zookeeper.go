package zookeeper

import (
	"log"
	"net"
	"time"

	"github.com/go-zookeeper/zk"
)

func simple() {
	conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}

	path := "/miya"
	var emptyData []byte

	if exists, _, err := conn.Exists(path); err == nil && exists {
		_ = conn.Delete(path, -1)
	}

	if path, err = conn.Create(path, emptyData, 0, zk.WorldACL(zk.PermAll)); err != nil {
		panic(err)
	}

	exists, _, err := conn.Exists(path)
	if err != nil || !exists {
		panic(err)
	}

	if _, err := conn.Set(path, []byte("hello miya"), -1); err != nil {
		panic(err)
	}

	if data, _, err := conn.Get(path); err != nil || string(data) != "hello miya" {
		panic(err)
	}

	//if err := conn.Delete(path, -1); err != nil {
	//	panic(err)
	//}
}

func connectBasic() {
	conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	children, _, err := conn.Children("/")
	if err != nil {
		panic(err)
	}

	log.Printf("children: %v", children)
}

func connectAdvance() {
	conn, events, err := zk.Connect([]string{"127.0.0.1"},
		time.Second*10,
		zk.WithDialer(net.DialTimeout),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		connected := false

		select {
		case event := <-events:
			log.Printf("received event from zk: %v", event)
			if event.State == zk.StateHasSession {
				log.Printf("session established")
				connected = true
			}
		}

		if connected {
			break
		}
	}

	// add auth
	err = conn.AddAuth("digest", []byte("root:12345"))
	if err != nil {
		panic(err)
	}

	// ls /go
	children, _, err := conn.Children("/go")
	if err != nil {
		panic(err)
	}
	log.Printf("children: %v", children)
}
