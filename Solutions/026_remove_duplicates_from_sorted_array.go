package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	pos := 0

	for _, num := range nums {
		if num > nums[pos] {
			nums[pos+1] = num
			pos++
		}
	}

	return pos + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := removeDuplicates(nums)
	fmt.Printf("New length: %d\n", k)
	fmt.Printf("Modified array: %v\n", nums[:k])
}
