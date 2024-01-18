package main

import (
	"fmt"
	"sync"
	"time"
)

var s sync.Mutex
var cnt = 0
var buy_cnt = 0
var provide_cnt = 0

func buy() {
	for {
		s.Lock()
		if cnt > 0 {
			cnt -= 1
			buy_cnt += 1
		}
		s.Unlock()
	}
}

func provide() {
	for {
		s.Lock()
		cnt += 1
		provide_cnt += 1
		s.Unlock()
	}
}

func main() {

	go buy()
	go provide()
	for {
		select {
		case <-time.Tick(time.Second * 1):
			s.Lock()
			fmt.Printf("cnt: %d, buy_cnt: %d, provide_cnt: %d\n", cnt, buy_cnt, provide_cnt)
			s.Unlock()
		}
	}
}
