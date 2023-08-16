package main

func permute(nums []int) [][]int {
	reslt := [][]int{}
	if len(nums) == 1 {
		return append([][]int{}, nums)
	}
	for i := 0; i < len(nums); i++ {
		n := nums[0]
		hoge := nums[1:len(nums)]
		perms := permute(hoge)
		for _, pv := range perms {
			pv = append(pv, n)
		}
		reslt = append(reslt, perms...)
		nums = append(nums, n)
	}
	return reslt
}
