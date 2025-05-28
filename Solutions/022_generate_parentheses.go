package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var res []string
	backtrack(n, n, "", &res)
	return res
}

func backtrack(open, close int, current string, res *[]string) {
	// If no more parentheses to add, the current string is complete.
	if open == 0 && close == 0 {
		*res = append(*res, current)
		return
	}
	// If we can add an opening parenthesis, do it.
	if open > 0 {
		backtrack(open-1, close, current+"(", res)
	}
	// Only add a closing parenthesis if it won't break the balance.
	if close > open {
		backtrack(open, close-1, current+")", res)
	}
}

func main() {
	n := 5
	result := generateParenthesis(n)
	fmt.Println(result)
}
