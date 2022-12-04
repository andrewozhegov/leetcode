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
	// fmt.Println(nums)
	min_diff := math.MaxInt
	min_diff_i := 0

	sums := make([]int, len(nums))
	for i, num := range nums {
		for j, _ := range nums[i:] {
			sums[i+j] += num
		}
	}
	fmt.Println(sums)
	all_sum := sums[len(sums)-1]

	for i, _ := range nums {
		// first_diff := sums[i] / (i + 1)
		last_diff := 0
		if len(sums) != i+1 {
			last_diff = (all_sum - sums[i]) / (len(sums) - i - 1)
		}
		// fmt.Println(" * ", last_diff)
		// fmt.Println("Last = ", last_diff)
		avg_diff := abs((sums[i] / (i + 1)) - last_diff)
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
