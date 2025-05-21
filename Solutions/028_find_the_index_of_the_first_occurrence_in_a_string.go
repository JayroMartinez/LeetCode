package main

import (
	"fmt"
)

func strStr(haystack, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}
	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "sadbutsad"
	needle := "but"

	res := strStr(haystack, needle)
	fmt.Printf("Position: %v\n", res)
}
