package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	cur := newHead
	l1cur := l1
	l2cur := l2
	curry := 0

	for l1cur != nil || l2cur != nil {
		var sum int
		if l1cur != nil {
			sum += l1cur.Val
			l1cur = l1cur.Next
		}
		if l2cur != nil {
			sum += l2cur.Val
			l2cur = l2cur.Next
		}
		add := (sum + curry) % 10
		nextNode := &ListNode{Val: add}
		cur.Next = nextNode
		cur = cur.Next
		curry = (sum + curry) / 10
	}
	if curry == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return newHead.Next
}
