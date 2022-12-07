package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) Insert(val int) *TreeNode {
	current := root
	previous := root

	if root == nil {
		root = &TreeNode{val, nil, nil}
		return root
	}

	for current != nil {
		previous = current
		if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	if val < previous.Val {
		previous.Left = &TreeNode{val, nil, nil}
	} else {
		previous.Right = &TreeNode{val, nil, nil}
	}

	return root
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var count int = 0
	if root.Left != nil {
		count += rangeSumBST(root.Left, low, high)
	}
	if root.Right != nil {
		count += rangeSumBST(root.Right, low, high)
	}
	if root.Val >= low && root.Val <= high {
		count += root.Val
	}
	return count
}

func main() {

	root := &TreeNode{10, nil, nil}
	for _, val := range []int{5, 15, 3, 7, 18} {
		root.Insert(val)
	}

	fmt.Println(rangeSumBST(root, 7, 15))
}
