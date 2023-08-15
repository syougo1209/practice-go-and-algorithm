package main

import "container/heap"

type IntHeap []int

type TimeControl struct {
	startTime int
	freq      int
}

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, v := range tasks {
		freq[v-'A']++
	}
	arr := []int{}
	for i := 0; i < 26; i++ {
		if freq[i] != 0 {
			arr = append(arr, (-1)*freq[i])
		}
	}
	maxHeap := IntHeap(arr)
	heap.Init(&maxHeap)
	queue := []TimeControl{}
	time := 0
	for len(maxHeap) != 0 || len(queue) != 0 {
		time++
		if len(maxHeap) != 0 {
			v, _ := heap.Pop(&maxHeap).(int)
			if v+1 != 0 {
				queue = append(queue, TimeControl{startTime: time + n, freq: v + 1})
			}
		}
		if len(queue) != 0 && queue[0].startTime == time {
			heap.Push(&maxHeap, queue[0].freq)
			tmp := len(queue)
			queue = queue[1:tmp]
		}
	}
	return time
}
