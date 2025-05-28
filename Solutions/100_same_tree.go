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

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	q := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	result := isSameTree(p, q)
	fmt.Println(result)
}
