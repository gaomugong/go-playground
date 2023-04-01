package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("exit from done")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// stop the ticker after 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

// Tick at 2023-04-01 19:52:23.78669 +0800 CST m=+0.500230584
// Tick at 2023-04-01 19:52:24.287595 +0800 CST m=+1.001137792
// Tick at 2023-04-01 19:52:24.787622 +0800 CST m=+1.501166376
// Ticker stopped
// exit from done
