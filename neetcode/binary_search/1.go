package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 12))
}

func search(nums []int, target int) int {

	var binarySearch func(l int, r int) int
	binarySearch = func(l int, r int) int {
		if l > r {
			return -1
		}
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return binarySearch(l, mid)
		} else {
			return binarySearch(mid+1, r)
		}
	}
	return binarySearch(0, len(nums))
}
