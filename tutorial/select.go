package main

import (
	"fmt"
	"time"
)

func main() {
	// select across two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// each channel receive a value after some seconds
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "second"
	}()

	// use select to await both of these values simultaneously and print
	// if read nothing in 3s, exit for timeout
	for i := 0; i < 3; i++ {
		fmt.Println("use select to await both of these values")
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
		}
	}

}
