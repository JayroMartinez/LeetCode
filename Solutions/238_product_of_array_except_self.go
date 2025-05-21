package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {

	length := len(nums)
	answer := make([]int, length)

	if length < 1 {

		return answer
	} else if length == 1 {
		answer[0] = 0
		return answer

	} else {

		aux := make([]int, length)

		aux[0] = 1

		for i := 1; i < length; i++ {
			aux[i] = aux[i-1] * nums[i-1]
		}

		answer[length-1] = aux[length-1]

		aux2 := 1
		for j := length - 2; j >= 0; j-- {
			aux2 = aux2 * nums[j+1]
			answer[j] = aux2 * aux[j]
		}
	}

	return answer

}

func main() {
	nums := []int{1, 2, 3, 4}

	res := productExceptSelf(nums)
	fmt.Printf("Result: %v\n", res)
}
