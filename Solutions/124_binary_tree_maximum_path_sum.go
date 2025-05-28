package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	// maxGain returns the maximum gain from this node downward.
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// Recursively get the maximum gain from left and right subtrees.
		// We discard negative gains by comparing with 0.
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// Price of the current node is the value plus both gains.
		currentPathSum := node.Val + leftGain + rightGain

		// Update global maximum if needed.
		if currentPathSum > maxSum {
			maxSum = currentPathSum
		}

		// Return the maximum gain the parent node could obtain by including this node.
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

// Helper function to get maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{-10,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}

	result := maxPathSum(root)
	fmt.Println(result)
}
