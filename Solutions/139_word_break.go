package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {

	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	memo := make(map[string]bool)
	return canBreak(s, wordSet, memo)
}

func canBreak(s string, wordSet map[string]bool, memo map[string]bool) bool {
	if s == "" {
		return true
	}

	if val, exists := memo[s]; exists {
		return val
	}

	for word := range wordSet {

		if len(s) >= len(word) && s[:len(word)] == word {

			if canBreak(s[len(word):], wordSet, memo) {
				memo[s] = true
				return true
			}
		}
	}
	memo[s] = false
	return false
}

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	result := wordBreak(s, wordDict)
	fmt.Println(result)
}
