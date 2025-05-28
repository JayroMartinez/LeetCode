package main

import (
	"fmt"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n)
	}

	for t := 1; t <= k; t++ {

		maxDiff := -prices[0]
		for d := 1; d < n; d++ {

			dp[t][d] = max(dp[t][d-1], prices[d]+maxDiff)

			maxDiff = max(maxDiff, dp[t-1][d]-prices[d])
		}
	}

	return dp[k][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	result := maxProfit(k, prices)
	fmt.Println(result)
}
