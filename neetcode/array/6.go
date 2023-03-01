package main

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	prefix := 1
	for i, v := range nums {
		ans[i] = prefix
		prefix *= v
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * postfix
		postfix *= ans[i]
	}
	return ans
}
