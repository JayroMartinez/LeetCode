package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// If the current number is greater than 0, break out
		// because we won't be able to find a triplet summing to 0.
		if nums[i] > 0 {
			break
		}
		// Skip duplicate values for i.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// Skip duplicate values for left pointer.
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicate values for right pointer.
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println("Result:", result)
}
