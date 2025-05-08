package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	for left := 0; left < len(nums); left++ {
		for right := left + 1; right < len(nums); right++ {
			if nums[left]+nums[right] == target {
				return []int{left, right}
			}
		}
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("Result:", result)
}
