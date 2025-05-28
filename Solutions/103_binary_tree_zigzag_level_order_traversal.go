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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {

		levelSize := len(queue)
		var level []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		if leftToRight {
			res = append(res, level)
		} else {
			var inverseLevel []int

			for it := len(level) - 1; it >= 0; it-- {
				inverseLevel = append(inverseLevel, level[it])
			}

			res = append(res, inverseLevel)
		}

		leftToRight = !leftToRight
	}

	return res
}

func main() {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}

	result := zigzagLevelOrder(root)
	fmt.Println(result)
}
