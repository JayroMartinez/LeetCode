package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {

	firstStr := strs[0]

	if len(strs) == 1 {
		return firstStr
	} else {

		finished := false

		for i := 0; i < len(firstStr); i++ { // iter over positions
			j := 1
			for j < len(strs) && !finished { // iter over strings

				if i >= len(strs[j]) || strs[j][i] != firstStr[i] {

					finished = true
					return firstStr[:i]
				}
				j++
			}
		}

	}

	return firstStr

}

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println("Result:", result)
}
