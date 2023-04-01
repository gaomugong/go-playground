package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	// block on the timer's channel
	<-timer1.C
	fmt.Println("Timer1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()
	// cancel timer2 before it fires
	if stop2 := timer2.Stop(); stop2 {
		fmt.Println("Timer2 stoped")
	}

	// Give timer2 enough time to fire
	time.Sleep(2 * time.Second)
}

//Timer1 fired
//Timer2 stoped
