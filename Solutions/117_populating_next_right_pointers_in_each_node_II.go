package main

import (
	"fmt"
)

// Definition for a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// If there's a left child, set its Next pointer.
	if root.Left != nil {
		if root.Right != nil {
			// If both children exist, left.Next points directly to right.
			root.Left.Next = root.Right
		} else {
			// Otherwise, find the next available child from root.Next.
			root.Left.Next = findNextChild(root.Next)
		}
	}

	// If there's a right child, set its Next pointer.
	if root.Right != nil {
		root.Right.Next = findNextChild(root.Next)
	}

	// Process the right subtree first so that when we process the left subtree,
	// the next pointers for nodes to the right are already established.
	connect(root.Right)
	connect(root.Left)

	return root
}

// findNextChild traverses the next pointers to find the first node that has a child.
func findNextChild(node *Node) *Node {
	for node != nil {
		if node.Left != nil {
			return node.Left
		}
		if node.Right != nil {
			return node.Right
		}
		node = node.Next
	}
	return nil
}

func printLevels(root *Node) {
	for levelStart := root; levelStart != nil; {
		for n := levelStart; n != nil; n = n.Next {
			fmt.Print(n.Val, " ")
		}
		fmt.Println("#") // fin de capa

		// buscar el primer nodo de la siguiente capa
		if levelStart.Left != nil {
			levelStart = levelStart.Left
		} else if levelStart.Right != nil {
			levelStart = levelStart.Right
		} else {
			levelStart = findNextChild(levelStart.Next)
		}
	}
}

func main() {
	root := &Node{1, &Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil}, &Node{3, &Node{6, nil, nil, nil}, &Node{7, nil, nil, nil}, nil}, nil}

	result := connect(root)
	printLevels(result)
}
