package main

type Node struct {
	Val  int
	key  int
	Next *Node
	Prev *Node
}
type LRUCache struct {
	head    *Node
	last    *Node
	NodeMap map[int]*Node
	Cap     int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{Cap: capacity}
	l.NodeMap = make(map[int]*Node)
	l.head = new(Node)
	l.last = new(Node)
	l.head.Next = l.last
	l.last.Prev = l.head
	return l
}
func (this *LRUCache) Remove(node *Node) {
	prev, nxt := node.Prev, node.Next
	prev.Next, nxt.Prev = nxt, prev
}

func (this *LRUCache) Insert(node *Node) {
	prev, nxt := this.last.Prev, this.last
	nxt.Prev = node
	prev.Next = nxt.Prev
	node.Next, node.Prev = nxt, prev
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.NodeMap[key]; ok {
		this.Remove(this.NodeMap[key])
		this.Insert(this.NodeMap[key])
		return this.NodeMap[key].Val
	} else {
		return -1
	}
}
func (this *LRUCache) Put(key int, value int) {
	node, ok := this.NodeMap[key]
	if ok {
		node.Val = value
		return
	}
	this.NodeMap[key] = &Node{Val: value, key: key}
	this.Insert(this.NodeMap[key])
	if len(this.NodeMap) > this.Cap {
		lru := this.head.Next
		this.Remove(lru)
		delete(this.NodeMap, lru.key)
	}
}
