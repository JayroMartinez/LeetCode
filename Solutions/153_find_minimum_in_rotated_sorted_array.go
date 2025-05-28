package main

import (
	"fmt"
	"sort"
)

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else {
		sort.Ints(nums)
	}

	return nums[0]
}

func main() {
	nums := []int{4, 2, 1, 3}
	result := findMin(nums)
	fmt.Println(result)
}
