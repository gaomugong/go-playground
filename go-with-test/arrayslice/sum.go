package arrayslice

func Sum(nums [5]int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
