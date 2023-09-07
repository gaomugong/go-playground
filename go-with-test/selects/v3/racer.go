package v1

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a string, b string) (winner string, err error) {
	return NewRacer(a, b, tenSecondTimeout)
}

// NewRacer select 中使用 ping 为两个 URL 设置两个 channel。
// 无论哪个先写入其 channel 都会使 select 里的代码先被执行，这会导致那个 URL 先被返回
func NewRacer(a string, b string, timeout time.Duration) (winner string, err error) {
	// select 则允许你同时在 多个 channel 等待。第一个发送值的 channel「胜出」，case 中的代码会被执行
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout(%s) waiting for %s and %s", timeout, a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		_, _ = http.Get(url)
		ch <- true
	}()

	return ch
}
