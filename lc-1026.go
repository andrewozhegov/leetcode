package main

import (
	"fmt"
	"strconv"
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
		// fmt.Println("V ", val, " [", i, "]", " -> ", getDepth(i+2))
		root.fill(val, getDepth(i+2))
	}
}

func (root *TreeNode) binaryTreePaths() []string {
	output := make([]string, 0)

	binaryTreePathsUtil(root, "", &output)
	return output
}

func binaryTreePathsUtil(root *TreeNode, curr string, output *[]string) {
	if root == nil {
		return
	}

	valString := strconv.Itoa(root.Val)
	if curr == "" {
		curr = valString
	} else {
		curr = curr + "->" + valString
	}
	if root.Left == nil && root.Right == nil {
		*output = append(*output, curr)
		return
	}

	binaryTreePathsUtil(root.Left, curr, output)
	binaryTreePathsUtil(root.Right, curr, output)

}

func findDiff(a []int) int {
	min, max := a[0], a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return max - min
}

func findAncestorDiff(root *TreeNode, path []int, diff *int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		tmp := findDiff(path)
		if *diff < tmp {
			*diff = tmp
		}
		return
	}
	if root.Left != nil {
		findAncestorDiff(root.Left, path, diff)
	}
	if root.Right != nil {
		findAncestorDiff(root.Right, path, diff)
	}
}

func maxAncestorDiff(root *TreeNode) int {
	diff := 0
	findAncestorDiff(root, []int{}, &diff)
	return diff
}

func main() {

	root := &TreeNode{8, nil, nil}
	vals := []int{3, 10, 1, 6, 1, 14, 1, 1, 4, 7, 13}
	root.fillTree(vals)

	fmt.Println(maxAncestorDiff(root))
}
