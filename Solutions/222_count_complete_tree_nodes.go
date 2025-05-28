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

func countNodes(root *TreeNode) int {

	result := 0

	if root == nil {
		return 0
	} else {
		result++
		result += countNodes(root.Left)
		result += countNodes(root.Right)
	}

	return result
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3,
			&TreeNode{6, nil, nil},
			nil}}

	result := countNodes(root)
	fmt.Println(result)
}
