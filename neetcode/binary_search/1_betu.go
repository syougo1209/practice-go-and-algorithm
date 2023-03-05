package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, -1))
}

func search(nums []int, target int) int {
	ans := -1
	l := 0
	r := len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			ans = mid
			break
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}
