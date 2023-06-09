package main

import "fmt"

func main() {
	// create empty map use `make`
	m := make(map[string]int)

	// set key/value pair
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))
	delete(m, "k2")
	fmt.Println("map: ", m)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`
	k2, ok := m["k2"]
	// k2, ok := m["k1"]
	fmt.Printf("k2: %d, ok: %v\n", k2, ok)

	// declare and initialize a new map in the same line.
	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map: ", n)
}
