package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	// We'll store the total sum of values per level and how many nodes per level.
	var totals []float64
	var counts []int

	// walk visits every node and tracks its level.
	var walk func(node *TreeNode, level int)
	walk = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// If this is a new level, add new entries.
		if level >= len(totals) {
			totals = append(totals, float64(node.Val))
			counts = append(counts, 1)
		} else {
			totals[level] += float64(node.Val)
			counts[level]++
		}
		walk(node.Left, level+1)
		walk(node.Right, level+1)
	}

	walk(root, 0)

	// Now, calculate the average at each level.
	averages := make([]float64, len(totals))
	for i := 0; i < len(totals); i++ {
		averages[i] = totals[i] / float64(counts[i])
	}
	return averages
}

func main() {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}

	result := averageOfLevels(root)
	fmt.Println(result)
}
