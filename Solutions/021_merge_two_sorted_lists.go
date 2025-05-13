package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // dummy node to ease list construction
	current := dummy

	// Merge while both lists have nodes.
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	// Append any remaining nodes.
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func main() {
	list1 := &ListNode{1, &ListNode{
		2, &ListNode{
			4, nil}}}
	list2 := &ListNode{1, &ListNode{
		3, &ListNode{
			4, nil}}}

	result := mergeTwoLists(list1, list2)

	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
	fmt.Println()
}
