package main

import (
	"fmt"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	min_diff, min_diff_i := math.MaxInt, 0

	prefix, suffix := 0, 0
	for _, num := range nums {
		suffix += num
	}

	for i, _ := range nums {
		prefix += nums[i]
		suffix -= nums[i]
		last_diff := 0
		if n != i+1 {
			last_diff = suffix / (n - i - 1)
		}
		avg_diff := abs(prefix/(i+1) - last_diff)
		if avg_diff < min_diff {
			min_diff = avg_diff
			min_diff_i = i
		}
	}
	return min_diff_i
}

func main() {
	data := []int{2, 5, 3, 9, 5, 3}
	fmt.Println(minimumAverageDifference(data))
}
