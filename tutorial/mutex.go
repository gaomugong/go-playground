package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps, writeOps uint64

	// start 100 gr to execute repeated read and sleep 1ms
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)

				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// start 10 gr to simulate writes using the same patterns we did for reads
	for r := 0; r < 10; r++ {
		go func() {
			for {
				key := rand.Intn(5)

				mutex.Lock()
				state[key] = rand.Intn(100)
				mutex.Unlock()

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// for w := 0; w < 10; w++ {
	// 	go func() {
	// 		for {
	// 			mutex.Lock()
	// 			fmt.Println("state: ", state)
	// 			mutex.Unlock()
	// 			time.Sleep(200 * time.Millisecond)
	// 		}
	// 	}()
	// }

	// let the 10 gr work on `state` and `mutex` for 1s
	time.Sleep(time.Second)

	// take and report final operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Printf("readOps: %d, writeOps: %d\n", readOpsFinal, writeOpsFinal)

	// lock and print state
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}
