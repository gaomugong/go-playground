package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{3, 2, 4, 5, 1, 8}
	strs := []string{"c", "f", "b", "a"}

	sort.Strings(strs)
	fmt.Println("strs: ", strs)

	sort.Ints(ints)
	fmt.Println("ints: ", ints, sort.IntsAreSorted(ints))
}
