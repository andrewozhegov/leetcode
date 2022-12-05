package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	head := &ListNode{7, nil}
	var data *ListNode = head
	for val := range []int{1, 2, 3, 4, 5, 6} {
		data.Next = &ListNode{val, nil}
		data = data.Next
	}
	fmt.Println(middleNode(data).Val)
}
