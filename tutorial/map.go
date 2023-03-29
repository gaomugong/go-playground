package main

import "fmt"

func main() {
	// create empty map use `make`
	m := make(map[string]int)

	// set key/value pair
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

}
