package main

func longestConsecutive(nums []int) int {
	hashMap := []bool{}
	for _, v := range nums {
		hashMap[v] = true
	}
	max := 0
	cnt := 1
	for num := range hashMap {
		if ok := hashMap[num-1]; ok {
			cnt++
		} else {
			cnt = 1
		}
		max = maxim(max, cnt)
	}
	return max
}

func maxim(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
