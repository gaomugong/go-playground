package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go f("first goroutine")
	go f("second goroutine")

	// start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 8)
	fmt.Println("done")
}
