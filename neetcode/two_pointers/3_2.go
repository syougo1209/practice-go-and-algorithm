package main

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i1 := 0; i1 < len(nums); i1++ {
		i2 := i1 + 1
		i3 := len(nums) - 1
		for i3 > i2 {
			n1 := nums[i1]
			n2 := nums[i2]
			n3 := nums[i3]
			sum := n1 + n2 + n3
			if sum == 0 {
				res = append(res, []int{n1, n2, n3})
				for i2 < len(nums) && n2 == nums[i2] {
					i2++
				}
				for n3 == nums[i3] && i3 >= 0 {
					i3--
				}
			} else {
				if sum > 0 {
					i3--
				} else {
					i2++
				}
			}
		}
	}
	return res
}
