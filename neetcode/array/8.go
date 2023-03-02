func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}

	ans := 0
	for i, v := range nums {
		if set[v-1] {
			continue
		}
		seq := 0
		temp := v
		for set[temp] {
			seq++
			temp++
		}
		if ans < seq {
			ans = seq
		}
	}
	return ans
}
