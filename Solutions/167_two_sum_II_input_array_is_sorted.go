package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {

	length := len(numbers)

	left := 0
	right := 1

	found := false

	for !found {
		if numbers[left]+numbers[right] == target {
			found = true
		} else if (right == length-1 && left < right-1) || numbers[left]+numbers[right] > target {
			left++
			right = left + 1
		} else if numbers[left]+numbers[right] < target {
			right++
		}

	}
	return []int{left + 1, right + 1}

}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(numbers, target)
	fmt.Println(result)
}
