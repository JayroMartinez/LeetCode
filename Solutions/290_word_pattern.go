package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {

	var aux map[string]string
	aux = make(map[string]string)

	var aux2 map[string]string
	aux2 = make(map[string]string)

	sep := strings.Fields(s)

	if len(sep) != len(pattern) {
		fmt.Println("Sizez do not match")
		return false
	}

	for it, word := range sep {

		if aux[word] == "" && aux2[string(pattern[it])] == "" {

			aux[word] = string(pattern[it])
			aux2[string(pattern[it])] = word

		} else if aux[word] != string(pattern[it]) || aux2[string(pattern[it])] != word {

			return false
		}
	}

	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"

	result := wordPattern(pattern, s)
	fmt.Println(result)
}
