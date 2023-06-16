package main

func containsDuplicate(nums []int) bool {
	dup := map[int]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		dup[num]++
		if dup[num] > 1 {
			return true
		}
	}
	return false
}
