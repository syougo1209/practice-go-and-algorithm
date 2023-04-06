package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	right := isSameTree(p.Right, q.Right)
	left := isSameTree(p.Left, q.Left)
	if right && left && p.Val == q.Val {
		return true
	} else {
		return false
	}
}
