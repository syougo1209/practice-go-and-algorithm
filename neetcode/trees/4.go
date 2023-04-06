package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	isB := true
	maxDepth(root, &isB)
	return isB
}
func maxDepth(node *TreeNode, isB *bool) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.Left, isB)
	right := maxDepth(node.Right, isB)
	if left > right+1 || right > left+1 {
		*isB = false
	}
	return maxim(left, right)
}
func maxim(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
