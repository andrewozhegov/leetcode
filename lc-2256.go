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

func abs_diff(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	for _, num := range nums {
		result += num
	}
	return result / len(nums)
}

func minimumAverageDifference(nums []int) int {
	min_diff := math.MaxInt
	min_diff_i := 0
	for i, _ := range nums {
		avg_diff := abs(abs_diff(nums[:i+1]...) - abs_diff(nums[i+1:]...))
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
