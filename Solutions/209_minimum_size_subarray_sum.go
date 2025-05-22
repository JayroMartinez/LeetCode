package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	left, sum, minLen := 0, 0, len(nums)+1

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	result := minSubArrayLen(target, nums)
	fmt.Println(result)
}
