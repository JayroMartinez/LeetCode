package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// need stores the frequency of each character required.
	need := make(map[rune]int)
	for _, ch := range t {
		need[ch]++
	}

	// missing: total number of characters (including duplicates) needed.
	missing := len(t)

	sRunes := []rune(s) // work with runes for Unicode safety
	left, start, minLen := 0, 0, len(sRunes)+1

	// Expand the window with right pointer.
	for right := 0; right < len(sRunes); right++ {
		// If current character is needed, reduce missing count.
		if count, exists := need[sRunes[right]]; exists {
			if count > 0 {
				missing--
			}
			need[sRunes[right]]--
		}

		// When we have all required characters, try shrinking the window.
		for missing == 0 {
			// Update minimum window if the current one is smaller.
			if right-left+1 < minLen {
				start = left
				minLen = right - left + 1
			}

			// Try to move left pointer ahead to shrink the window.
			if _, exists := need[sRunes[left]]; exists {
				need[sRunes[left]]++
				// If this character becomes required, increase missing.
				if need[sRunes[left]] > 0 {
					missing++
				}
			}
			left++
		}
	}

	if minLen > len(sRunes) {
		return ""
	}
	return string(sRunes[start : start+minLen])
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	result := minWindow(s, t)
	fmt.Println(result)
}
