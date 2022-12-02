package main

import (
	"fmt"
	"strings"
)

func halvesAreAlike(s string) bool {
	var counter uint = 0
	for i := 0; i < len(s)/2; i++ {
		if strings.ContainsRune("aeiouAEIOU", rune(s[i])) {
			counter++
		}
		if strings.ContainsRune("aeiouAEIOU", rune(s[i+len(s)/2])) {
			counter--
		}
	}
	return 0 == counter
}

func main() {
	fmt.Println(halvesAreAlike("book"))
}
