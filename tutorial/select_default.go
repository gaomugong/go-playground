/*
这段代码是Go语言中关于使用select实现非阻塞发送和接收的示例。
在Go语言中，通道的基本发送和接收操作是阻塞的，即当通道未准备好接收或发送新值时，程序将被阻塞等待，直到有新值被接收或发送完成。
但是，我们可以使用select语句和default子句来实现非阻塞发送和接收操作，从而避免程序被阻塞等待。
*/
package main

import "fmt"

func main() {
	messages := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		// 如果没有任何通道准备好接收新值，则select语句将立即选择default子句执行
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("sent message", msg)
	case sig := <-signals:
		fmt.Println("sent message", sig)
	default:
		// 如果多个通道都没有准备好接收新值，则select语句将立即选择default子句执行
		fmt.Println("no activity")
	}
}
