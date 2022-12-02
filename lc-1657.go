package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	const length = len("abcdefghijklmnopqrstuvwxyz")
	const rune_inc = 97

	count1, count2 := make([]int, length), make([]int, length)
	for _, rune := range word1 {
		count1[rune-rune_inc]++
	}
	for _, rune := range word2 {
		count2[rune-rune_inc]++
	}

	for i := 0; i < length; i++ {
		if (count1[i] == count2[i]) || (count1[i] > 0 && count2[i] > 0) {
			continue
		}
		return false
	}

	sort.Ints(count1)
	sort.Ints(count2)

	for i := 0; i < length; i++ {
		if count1[i] != count2[i] {
			return false
		}
	}
	return true
}

func main() {
	// fmt.Println(closeStrings("abca", "bcaa"))
	fmt.Println(closeStrings("abbzzca", "babzzcz"))
}
