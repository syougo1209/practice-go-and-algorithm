package main

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
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	laterFirst := slow.Next
	slow.Next = nil
	last := reverse(laterFirst)
	tmp := head
	dummy := new(ListNode)
	cur := dummy
	for tmp != nil {
		cur.Next = tmp
		tmp = tmp.Next
		cur = cur.Next
		if last != nil {
			cur.Next = last
			last = last.Next
			cur = cur.Next
		}
	}
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}
