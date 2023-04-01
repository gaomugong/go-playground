// In the previous example we used explicit locking with
// [mutexes](mutexes) to synchronize access to shared state
// across multiple goroutines. Another option is to use the
// built-in synchronization features of  goroutines and
// channels to achieve the same result. This channel-based
// approach aligns with Go's ideas of sharing memory by
// communicating and having each piece of data owned
// by exactly 1 goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 封装读写请求结构体，内部含有用于通信的channel
type readOp struct {
	key int
	// receive value from resp channel
	resp chan int
}

// writeOp is used to store data related to a write operation.
// It contains a key, value, and a response channel.
type writeOp struct {
	key int
	val int
	// receive write ops result from resp channel
	resp chan bool
}

func main() {
	var readOps, writeOps uint64
	// read and write channels which carry the r/w reqs
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// the gr which owns the `state` which handle w/r reqs from other gr and make response
	// 由一个gr保管中心化的state，通过借助channel与这个gr通信的方式来共享state
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case readReq := <-reads:
				readReq.resp <- state[readReq.key]
			case writeReq := <-writes:
				state[writeReq.key] = writeReq.val
				writeReq.resp <- true
			}
		}
	}()

	// start 100 gr to execute repeated read and sleep 1ms
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				readReq := readOp{rand.Intn(5), make(chan int)}
				reads <- readReq
				total += <-readReq.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// start 10 gr to simulate writes using the same patterns we did for reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				writeReq := writeOp{rand.Intn(5), rand.Intn(100), make(chan bool)}
				writes <- writeReq
				isOk := <-writeReq.resp
				if isOk {
					atomic.AddUint64(&writeOps, 1)
				}
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let the 10 gr work on `state` and `mutex` for 1s
	time.Sleep(time.Second)

	// take and report final operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Printf("readOps: %d, writeOps: %d\n", readOpsFinal, writeOpsFinal)
}
