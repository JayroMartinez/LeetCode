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

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left != nil {
		leftRes := isValidBST(root.Left) && maxTree(root.Left) < root.Val

		if !leftRes {
			return false
		}
	}

	if root.Right != nil {
		rightRes := isValidBST(root.Right) && minTree(root.Right) > root.Val

		if !rightRes {
			return false
		}
	}

	return true
}

func maxTree(root *TreeNode) int {

	maxVal := root.Val

	if root.Left != nil {
		auxL := maxTree(root.Left)
		if auxL > maxVal {
			maxVal = auxL
		}
	}

	if root.Right != nil {
		auxR := maxTree(root.Right)
		if auxR > maxVal {
			maxVal = auxR
		}
	}

	return maxVal
}

func minTree(root *TreeNode) int {
	minVal := root.Val

	if root.Left != nil {
		auxL := minTree(root.Left)
		if auxL < minVal {
			minVal = auxL
		}
	}

	if root.Right != nil {
		auxR := minTree(root.Right)
		if auxR < minVal {
			minVal = auxR
		}
	}

	return minVal

}

func main() {
	root := &TreeNode{3,
		&TreeNode{2, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{35, nil, nil}}}

	result := isValidBST(root)
	fmt.Println(result)
}
