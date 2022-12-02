package main

import (
	"fmt"
	"sort"
)

// If the key(values of the slice) is not equal
// to the already present value in new slice (list)
// then we append it. else we jump on another element.
func unique(slice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	runes1 := []rune(word1)
	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})

	runes2 := []rune(word2)
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	return string(unique(runes1)) == string(unique(runes2))
}

func main() {
	// fmt.Println(closeStrings("abca", "bcaa"))
	fmt.Println(closeStrings("abbzzca", "babzzcz"))
}
