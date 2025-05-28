package main

import (
	"fmt"
)

func hammingWeight(n int) int {

	count := 0
	for n != 0 {
		n &= n - 1 // clear the least significant bit set
		count++
	}
	return count
}

func main() {
	n := 2147483645
	result := hammingWeight(n)
	fmt.Println(result)
}
