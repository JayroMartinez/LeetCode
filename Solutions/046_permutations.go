package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var results [][]int

	var backtrack func(startIndex int)
	backtrack = func(startIndex int) {
		// Base case: if we've fixed all positions, record the permutation.
		if startIndex == len(nums)-1 {
			// Make a copy of nums and append to results.
			perm := make([]int, len(nums))
			copy(perm, nums)
			results = append(results, perm)
			return
		}

		// Try swapping the current index with each candidate index.
		for i := startIndex; i < len(nums); i++ {
			// Swap element at startIndex with element at i.
			nums[startIndex], nums[i] = nums[i], nums[startIndex]
			// Recurse with next index.
			backtrack(startIndex + 1)
			// Backtrack: swap back to restore the original order.
			nums[startIndex], nums[i] = nums[i], nums[startIndex]
		}
	}

	backtrack(0)
	return results
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println(result)
}
