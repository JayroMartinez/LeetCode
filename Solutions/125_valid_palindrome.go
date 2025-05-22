package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	lower := strings.ToLower(s)

	clean := removeNonAlfa(lower)

	start := 0
	end := len(clean) - 1

	for end >= start {
		if clean[start] == clean[end] {
			start++
			end--
		} else {
			return false
		}
	}

	return true
}

func removeNonAlfa(s string) string {
	result := ""

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || (ch >= '0' && ch <= '9') {
			result += string(ch)
		}
	}

	return result

}

func main() {
	s := "A man, a plan, a canal: Panama"

	result := isPalindrome(s)
	fmt.Println(result)
}
