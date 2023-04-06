package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	maxDepth(root, &max)
	return max
}

func maxDepth(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.Left, max)
	right := maxDepth(node.Right, max)
	*max = maxim(*max, left+right)
	return maxim(left, right) + 1
}

func maxim(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
