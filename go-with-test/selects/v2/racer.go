package v1

import (
	"net/http"
)

// Racer select 中使用 ping 为两个 URL 设置两个 channel。
// 无论哪个先写入其 channel 都会使 select 里的代码先被执行，这会导致那个 URL 先被返回
func Racer(a string, b string) (winner string) {
	// select 则允许你同时在 多个 channel 等待。第一个发送值的 channel「胜出」，case 中的代码会被执行
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
