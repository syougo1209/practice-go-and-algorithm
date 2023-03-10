package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	tmp := head
	for tmp != nil {
		newNode := &Node{Val: tmp.Val}
		nodeMap[tmp] = newNode
		tmp = tmp.Next
	}

	tmp = head
	cur := nodeMap[tmp]
	for tmp != nil {
		cur.Next = nodeMap[tmp.Next]
		cur.Random = nodeMap[tmp.Random]
		cur = cur.Next
		tmp = tmp.Next
	}
	return nodeMap[head]
}
