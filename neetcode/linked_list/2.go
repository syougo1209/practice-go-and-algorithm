package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := new(ListNode)
	dummy := head

	cur1 := l1
	cur2 := l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			dummy.Next = cur2
			cur2 = cur2.Next
			dummy = dummy.Next
		} else {
			dummy.Next = cur1
			cur1 = cur1.Next
			dummy = dummy.Next
		}
	}

	for cur1 != nil {
		dummy.Next = cur1
		cur1 = cur1.Next
		dummy = dummy.Next
	}
	for cur2 != nil {
		dummy.Next = cur2
		cur2 = cur2.Next
		dummy = dummy.Next
	}
	head = head.Next
	return head
}
