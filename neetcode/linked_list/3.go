package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	forth := &ListNode{Val: 4, Next: nil}
	third := &ListNode{Val: 3, Next: forth}
	second := &ListNode{Val: 2, Next: third}
	first := &ListNode{Val: 1, Next: second}
	head := first
	reorderList(head)
}
func reorderList(head *ListNode) {
	tmp := head

	arr := make([]*ListNode, 0)
	for tmp.Next != nil {
		arr = append(arr, tmp)
		tmp = tmp.Next
	}
	arr = append(arr, tmp)
	l := 0
	r := len(arr) - 1
	dummy := new(ListNode)
	cur := dummy
	for l <= r {
		cur.Next = arr[l]
		cur = cur.Next
		cur.Next = arr[r]
		cur = cur.Next
		l++
		r--
	}

	cur.Next = nil
	tmp = dummy.Next
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}
