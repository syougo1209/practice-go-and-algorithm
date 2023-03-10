package main

func hasCycle(head *ListNode) bool {
	cur := head
	fast := cur
	slow := cur
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
