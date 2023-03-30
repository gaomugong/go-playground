package main

import (
	"fmt"
	"sort"
)

type byLength []string

// Len 实现sort.Interface接口
func (s byLength) Len() int {
	return len(s)
}

// Less 按字符串长度排序
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	ints := []int{3, 2, 4, 5, 1, 8}
	strs := []string{"c", "f", "b", "a"}

	sort.Strings(strs)
	fmt.Println("strs: ", strs)

	sort.Ints(ints)
	fmt.Println("ints: ", ints, sort.IntsAreSorted(ints))

	animals := []string{"bear", "dog", "rabbit", "panda", "cat"}
	sort.Sort(byLength(animals))
	fmt.Println(animals)
}
