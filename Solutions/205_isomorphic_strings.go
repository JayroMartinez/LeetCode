package main

import (
	"fmt"
)

func isIsomorphic(s, t string) bool {
	// If the strings are of different lengths, they cannot be isomorphic.
	if len(s) != len(t) {
		return false
	}

	// Map from characters in s to characters in t.
	sToT := make(map[rune]rune)
	// Map from characters in t to characters in s.
	tToS := make(map[rune]rune)

	// Iterate over the characters.
	for i, chS := range s {
		chT := rune(t[i])
		// Check mapping from s to t.
		if mapped, ok := sToT[chS]; ok {
			if mapped != chT {
				return false
			}
		} else {
			sToT[chS] = chT
		}

		// Check reverse mapping from t to s.
		if mapped, ok := tToS[chT]; ok {
			if mapped != chS {
				return false
			}
		} else {
			tToS[chT] = chS
		}
	}

	return true
}

func main() {
	s := "egg"
	t := "add"

	result := isIsomorphic(s, t)
	fmt.Println(result)
}
