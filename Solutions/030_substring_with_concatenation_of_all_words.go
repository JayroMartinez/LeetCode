package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	// Edge cases.
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	totalWords := len(words)
	totalLen := wordLen * totalWords
	if len(s) < totalLen {
		return []int{}
	}

	// Build a frequency map for the words.
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	var res []int

	// Iterate over every possible starting index.
	for i := 0; i <= len(s)-totalLen; i++ {
		// For the window starting at i, create a map for the words seen.
		seen := make(map[string]int)
		valid := true

		// Divide the window into totalWords segments of length wordLen.
		for j := 0; j < totalWords; j++ {
			start := i + j*wordLen
			end := start + wordLen
			sub := s[start:end]
			// If the word is not in freq, the window is invalid.
			if count, ok := freq[sub]; !ok {
				valid = false
				break
			} else {
				seen[sub]++
				// If we've seen the word more times than required, window is invalid.
				if seen[sub] > count {
					valid = false
					break
				}
			}
		}

		if valid {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}

	result := findSubstring(s, words)
	fmt.Println(result)
}
