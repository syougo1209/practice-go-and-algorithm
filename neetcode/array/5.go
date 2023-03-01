package main

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	topKFrequent(nums, 2)
}
func topKFrequent(nums []int, k int) []int {
	//sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	freq := map[int]int{}
	for _, v := range nums {
		freq[v] += 1
	}
	countSlice := make([][]int, len(nums)+1)
	for i, v := range freq {
		countSlice[v] = append(countSlice[v], i)
	}
	res := []int{}
	for i := len(nums); i > 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			return res
		}
	}

	return res
}
