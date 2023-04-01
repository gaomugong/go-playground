package main

import "fmt"

func main() {
	// create a string channel
	messages := make(chan string)

	// send "ping" to channel messages in a goroutine
	go func() { messages <- "ping" }()

	// read "ping" from messages
	msg := <-messages
	fmt.Println(msg)

	// make a string channel buffering up to 2 values
	messages1 := make(chan string, 2)
	messages1 <- "buffered"
	messages1 <- "channel"

	fmt.Println(<-messages1)
	fmt.Println(<-messages1)
}
