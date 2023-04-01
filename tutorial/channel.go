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
}
