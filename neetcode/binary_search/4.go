package main

func findMin(nums []int) int {
	target := nums[0]
	l := 0
	r := len(nums)
	ans := nums[0]
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > target { //midがleft側の配列
			l = mid + 1
		} else { //midがright側の配列
			r = mid
			ans = min(ans, nums[mid])
		}
	}
	return ans
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
