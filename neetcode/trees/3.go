package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, max := maxsize(root)
	return max
}

func maxsize(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	left, ml := maxsize(root.Left)
	right, mr := maxsize(root.Right)
	max := maxim(left+right, ml)
	max = maxim(max, mr)

	return maxim(left, right) + 1, max
}

func maxim(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
