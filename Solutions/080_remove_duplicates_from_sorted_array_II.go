package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}

	pos := 2

	for iter := 2; iter < len(nums); iter++ {
		if nums[iter] != nums[pos-2] {
			nums[pos] = nums[iter]
			pos++
		}
	}

	return pos

}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}

	k := removeDuplicates(nums)
	fmt.Printf("New length: %d\n", k)
	fmt.Printf("Modified array: %v\n", nums[:k])
}
