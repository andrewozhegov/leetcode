package main

import (
	"fmt"
	"reflect"
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

func (root *TreeNode) fill(val int, dep int) *TreeNode {
	if dep == 0 {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		root.Left = &TreeNode{val, nil, nil}
		return root.Left
	}
	if root.Right == nil {
		root.Right = &TreeNode{val, nil, nil}
		return root.Right
	}

	left := root.Left.fill(val, dep-1)
	if left != nil {
		return left
	}
	right := root.Right.fill(val, dep-1)
	if right != nil {
		return right
	}

	return nil
}

func getDepth(index int) int {
	dep := 0
	for bin := 2; bin <= index; bin *= 2 {
		dep++
	}
	return dep
}

func (root *TreeNode) fillTree(values []int) {
	for i, val := range []int{5, 1, 6, 2, 9, 8, 7, 4} {
		fmt.Println("V ", val, " [", i, "]", " -> ", getDepth(i+2))
		root.fill(val, getDepth(i+2))
	}
}

func countLeaf(root *TreeNode) []int {
	count := []int{}
	if root.Left != nil {
		count = append(count, countLeaf(root.Left)...)
	}
	if root.Right != nil {
		count = append(count, countLeaf(root.Right)...)
	}
	if root.Left == nil && root.Right == nil {
		count = append(count, root.Val)
	}
	return count
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(countLeaf(root1), countLeaf(root2))
}

func main() {

	root1 := &TreeNode{3, nil, nil}
	vals1 := []int{5, 1, 6, 2, 9, 8, 7, 4}
	root1.fillTree(vals1)

	root2 := &TreeNode{3, nil, nil}
	vals2 := []int{5, 1, 6, 7, 4, 2, 9, 8}
	root2.fillTree(vals2)

	fmt.Println(leafSimilar(root1, root2))
}
