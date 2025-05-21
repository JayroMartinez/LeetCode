package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	h := 0

	sort.Ints(citations)
	length := len(citations)

	if length == 0 {
		return h
	} else {

		for iter := 0; iter < length; iter++ {
			if citations[length-iter-1] > iter {
				h++
			} else {
				return h
			}
		}
	}
	return h
}

func main() {
	citations := []int{3, 0, 6, 1, 5}

	prof := hIndex(citations)
	fmt.Printf("H-Index: %d\n", prof)
}
