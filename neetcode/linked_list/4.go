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
	fmt.Println(removeNthFromEnd(head, 1))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := head
	length := 0
	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	target := length - n + 1
	cur := 1
	tmp = head
	var prev *ListNode
	for cur != target && tmp != nil {
		cur++
		prev = tmp
		tmp = tmp.Next
	}
	if tmp.Next != nil && prev.Next != nil {
		prev.Next = tmp.Next
	} else if tmp.Next == nil {
		prev.Next = nil
	} else if prev == nil {
		head = tmp.Next
	}
	return head
}
