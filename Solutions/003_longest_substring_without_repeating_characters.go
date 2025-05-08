package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxLen := 0
	seen := make(map[rune]int)

	for i, ch := range s {
		if last, ok := seen[ch]; ok && last >= start {
			start = last + 1
		}
		seen[ch] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Println("Result:", result)
}
