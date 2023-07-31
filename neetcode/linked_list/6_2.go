package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head
	carry := 0
	l1cur := l1
	l2cur := l2
	for l1cur != nil || l2cur != nil {
		sum := carry
		if l1cur != nil {
			sum += l1cur.Val
			l1cur = l1cur.Next
		}

		if l2cur != nil {
			sum += l2cur.Val
			l2cur = l2cur.Next
		}
		node := &ListNode{Val: sum % 10}
		cur.Next = node
		carry = sum / 10
	}
	if carry > 0 {
		node := new(ListNode)
		node.Val = carry
		cur.Next = node
	}

	return head.Next
}
