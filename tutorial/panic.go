package main

import "fmt"

// This function panics.
func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from: ", err)
		}
	}()

	fmt.Println("starting the program...")
	//panic("a server error occurred")
	mayPanic()
	fmt.Println("this line will never be printed")
}
