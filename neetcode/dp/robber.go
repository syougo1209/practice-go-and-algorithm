package main

func maxi(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func rob(nums []int) int {
	one := 0
	two := nums[0]
	max := 0
	for i, v := range nums {
		if i == 0 {
			continue
		}
		tmp := two
		two = maxi(one, two+v)
		max = maxi(max, two)
		one = tmp
	}
	return max
}
