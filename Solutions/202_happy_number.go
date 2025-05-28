package main

import (
	"fmt"
)

func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}

func sumOfSquares(n int) int {
	total := 0
	for n > 0 {
		digit := n % 10
		total += digit * digit
		n /= 10
	}
	return total
}

func main() {
	n := 19

	result := isHappy(n)
	fmt.Println(result)
}
