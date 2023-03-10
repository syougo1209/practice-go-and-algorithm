package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	hashMap := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if hashMap[cur] {
			return true
		}
		hashMap[cur] = true
		cur = cur.Next
	}
	return false
}
