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

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := len(nums) / 2
	root := &TreeNode{Val: nums[midIndex]}
	root.Left = sortedArrayToBST(nums[:midIndex])
	root.Right = sortedArrayToBST(nums[midIndex+1:])
	return root
}

func toLevelOrder(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}
	out := []any{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, n.Val)
		if n.Left != nil || n.Right != nil {
			q = append(q, n.Left, n.Right)
		}
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	result := sortedArrayToBST(nums)
	fmt.Println(toLevelOrder(result))
}
