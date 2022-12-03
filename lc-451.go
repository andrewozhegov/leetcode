package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	c := make(map[rune]int)
	for _, char := range s {
		_, ok := c[char]
		if !ok {
			c[char] = 0
		}
		c[char]++
	}

	keys := make([]rune, 0, len(c))
	for key := range c {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return c[keys[i]] > c[keys[j]]
	})

	var ss string = ""
	for _, sym := range keys {
		for c[rune(sym)] > 0 {
			ss += string(sym)
			c[rune(sym)]--
		}
	}

	return ss
}

func main() {
	fmt.Println(frequencySort("Aabb"))
	// fmt.Println(frequencySort("jfhgjfhgjfhgghj"))
	// fmt.Println(frequencySort("tree"))
}
