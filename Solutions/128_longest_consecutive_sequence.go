package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	// Create a dummy node to simplify edge cases.
	dummy := &ListNode{Next: head}
	// groupPrev is the node right before the current group.
	groupPrev := dummy

	// Process groups until there are fewer than k nodes left.
	for {
		// Find the kth node from groupPrev.
		kth := getKthNode(groupPrev, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next // Node following the kth node.

		// Reverse the group. Start at groupPrev.Next and reverse nodes up to kth.
		prev := groupNext
		curr := groupPrev.Next
		for curr != groupNext {
			nextTemp := curr.Next
			curr.Next = prev
			prev = curr
			curr = nextTemp
		}

		// After reversal, groupPrev.Next is the kth node (new head of the group).
		// Save the start of the group (which becomes the new tail after reversal).
		newTail := groupPrev.Next
		groupPrev.Next = kth

		// Move groupPrev to the newTail for next iteration.
		groupPrev = newTail
	}
	return dummy.Next
}

// getKthNode returns the kth node from the current node, or nil if fewer than k nodes remain.
func getKthNode(current *ListNode, k int) *ListNode {
	for current != nil && k > 0 {
		current = current.Next
		k--
	}
	return current
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	result := reverseKGroup(head, 3)
	//fmt.Println(result)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
