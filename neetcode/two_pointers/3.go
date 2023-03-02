package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		end := len(nums) - 1
		mid := i + 1
		for mid < end {
			sum := nums[i] + nums[mid] + nums[end]

			if sum == 0 {
				end--
				ans = append(ans, []int{nums[i], nums[mid], nums[end]})
				for mid < end && nums[end] == nums[end+1] {
					end--
				}
			} else if sum > 0 {
				end--
			} else if sum < 0 {
				mid++
			}
		}
	}
	return ans
}
