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
	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
	} else if this.mins[len(this.mins)-1] > val {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, this.mins[len(this.mins)-1])
	}
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
