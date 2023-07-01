package main

type MinStack struct {
	nums []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
	mins[len]
}

func (this *MinStack) Pop() {
}

func (this *MinStack) Top() int {

}

func (this *MinStack) GetMin() int {

}
