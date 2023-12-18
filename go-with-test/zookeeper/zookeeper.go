package zookeeper

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-zookeeper/zk"
)

type ZK struct {
	conn *zk.Conn
}

func (z *ZK) connect(url string) (*zk.Conn, error) {
	conn, _, err := zk.Connect([]string{url}, time.Second*5)
	if err != nil {
		return nil, fmt.Errorf("new zk connection error: %w", err)
	}

	z.conn = conn
	return nil, nil
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

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}

	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
