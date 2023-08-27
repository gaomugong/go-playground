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
	queue := make(chan string, 2)
	queue <- "buffered"
	queue <- "channel"

	fmt.Println(<-queue)
	fmt.Println(<-queue)

	// 遍历时，如果channel 没有关闭，那么会一直等待下去，出现 deadlock 的错误；
	// 如果在遍历时channel已经关闭，那么在遍历完数据后自动退出遍历
	queue <- "one"
	queue <- "two"
	// fatal error: all goroutines are asleep - deadlock!
	close(queue)
	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
