package main

func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{}
	countSlice := map[int][]int{}
	for _, v := range nums {
		countMap[v] += 1
	}
	for i, v := range countMap {
		countSlice[v] = append(countSlice[v], i)
	}
	res := []int{}
	for i := len(countSlice); i > 0; i-- {
		res = append(res, countSlice[i]...)
		if k == len(res) {
			return res
		}
	}
	return res
}
