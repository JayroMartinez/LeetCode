package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	pos := 0 // Position to place the next non-val element

	for _, num := range nums {
		if num != val {
			nums[pos] = num // Place valid elements at the front
			pos++
		}
	}

	return pos // The new length of the modified array
}

func main() {
	nums := []int{3, 2, 2, 3, 4, 3}
	val := 3

	k := removeElement(nums, val)
	fmt.Printf("New length: %d\n", k)
	fmt.Printf("Modified array: %v\n", nums[:k])
}
