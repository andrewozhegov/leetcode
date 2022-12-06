package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	head_e := head.Next
	for odd != nil && even != nil {
		if odd.Next != nil {
			odd.Next = odd.Next.Next
		}
		if even.Next != nil {
			even.Next = even.Next.Next
		}
		if odd.Next == nil {
			odd.Next = head_e
			break
		}
		if even.Next == nil {
			odd.Next.Next = head_e
			break
		}
		odd = odd.Next
		even = even.Next
	}
	return head
}

func main() {
	head := &ListNode{0, nil}
	var data *ListNode = head
	for _, val := range [8]int{1, 2, 3, 4, 5, 6, 7, 8} {
		data.Next = &ListNode{val, nil}
		data = data.Next
	}
	for iter := head; iter != nil; iter = iter.Next {
		fmt.Print(iter.Val, " ")
	}
	fmt.Println()
	for iter := oddEvenList(head); iter != nil; iter = iter.Next {
		fmt.Print(iter.Val, " ")
	}
	fmt.Println()
}
