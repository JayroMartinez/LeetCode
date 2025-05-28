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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root} // Initialize the queue with the root node

	// While there are nodes to process:
	for len(queue) > 0 {
		levelSize := len(queue) // Number of nodes in the current level
		var level []int         // This will store the values of the current level

		// Process all nodes at this level:
		for i := 0; i < levelSize; i++ {
			// Pop the first node from the queue
			node := queue[0]
			queue = queue[1:]

			// Append the current node's value to the level slice
			level = append(level, node.Val)

			// Add the children to the queue for the next level:
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// Append the current level's values to the result
		res = append(res, level)
	}
	return res
}

func main() {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}

	result := levelOrder(root)
	fmt.Println(result)
}
