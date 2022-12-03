package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	count := make([]int, 128)
	for _, char := range s {
		count[char]++
	}
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return (count[r[i]] > count[r[j]]) ||
			(count[r[i]] == count[r[j]]) &&
				r[i] > r[j]
	})
	return string(r)
}

func main() {
	// fmt.Println(frequencySort("Aabb"))
	// fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("loveleetcode"))
}
