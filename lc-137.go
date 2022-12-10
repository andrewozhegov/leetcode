package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var one int = 0
	var two int = 0

	for i := 0; i < len(nums); i++ {
		one = (one ^ nums[i]) & ^two
		two = (two ^ nums[i]) & ^one
	}

	return one + two
}

func main() {
	inpu := []int{1, 1, 1, 2}
	fmt.Println(singleNumber(inpu))
}
