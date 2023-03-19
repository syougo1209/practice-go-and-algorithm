package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}
