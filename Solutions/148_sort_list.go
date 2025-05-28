package main

import (
	"fmt"
	"sort"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	var auxArr []int

	if head == nil {
		return nil
	} else {
		for head != nil {
			auxArr = append(auxArr, head.Val)
			head = head.Next
		}
	}

	sort.Ints(auxArr)

	dummy := &ListNode{}
	current := dummy

	for _, v := range auxArr {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return dummy.Next
}

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	result := sortList(head)
	for result != nil {
		fmt.Println(result.Val, " ")
		result = result.Next
	}
}
