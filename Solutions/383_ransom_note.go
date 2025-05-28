package main

import (
	"fmt"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {

	if len(magazine) < len(ransomNote) {
		return false
	} else {
		for _, letter := range ransomNote {
			if strings.Count(ransomNote, string(letter)) > strings.Count(magazine, string(letter)) {
				return false
			}
		}
	}

	return true
}

func main() {
	ransomNote := "aa"
	magazine := "aab"

	result := canConstruct(ransomNote, magazine)
	fmt.Println(result)
}
