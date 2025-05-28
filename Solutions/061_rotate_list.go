package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// Handle edge cases.
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Compute the length of the list and get the tail.
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// Normalize k: rotating by a multiple of length yields the original list.
	k = k % length
	if k == 0 {
		return head
	}

	// Connect tail to head, making it circular.
	tail.Next = head

	// Find the new tail: it is (length - k) nodes from the beginning.
	stepsToNewTail := length - k
	newTail := head
	for i := 0; i < stepsToNewTail-1; i++ {
		newTail = newTail.Next
	}

	// The node after newTail is the new head.
	newHead := newTail.Next
	// Break the circle.
	newTail.Next = nil

	return newHead
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	k := 4
	result := rotateRight(head, k)
	//fmt.Println(result)
	for result != nil {
		fmt.Println(result.Val, " ")
		result = result.Next
	}
}
