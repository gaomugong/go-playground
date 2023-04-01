package main

import "fmt"

// factorial is a recursive function that computes the factorial of an integer.
// The factorial of a given number n is defined as the product of all its numbers from 1 to n.
// If n is equal to 0, the factorial is equal to 1.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(3))
	fmt.Println(factorial(4))
}
