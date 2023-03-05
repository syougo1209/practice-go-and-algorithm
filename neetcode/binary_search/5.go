package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Println(search(nums, 0))
}

func search(nums []int, target int) int {
	first := nums[0]
	if target == first {
		return 0
	} else if target > first { //左を探す
		return searchLeft(nums, target)
	} else if target < first { //右を探す
		return searchRight(nums, target)
	}
	return -1
}
func searchLeft(nums []int, target int) int {
	l := 0
	r := len(nums)
	first := nums[0]
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < first {
			r = mid
		} else if nums[mid] >= first && nums[mid] < target {
			l = mid + 1
		} else if nums[mid] >= first && nums[mid] > target {
			r = mid
		}
	}
	return -1
}

func searchRight(nums []int, target int) int {
	l := 0
	r := len(nums)
	first := nums[0]
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > first {
			l = mid + 1
		} else if nums[mid] <= first && nums[mid] < target {
			l = mid + 1
		} else if nums[mid] <= first && nums[mid] > target {
			r = mid
		}
	}
	return -1
}
