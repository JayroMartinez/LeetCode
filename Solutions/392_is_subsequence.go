package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	init := 0

	for _, it := range s {

		index := strings.Index(t[init:], string(it))
		if index == -1 {
			return false
		}

		init += index + 1
	}

	return true

}

func main() {
	s := "abc"
	t := "ahbgdc"

	result := isSubsequence(s, t)
	fmt.Println(result)
}
