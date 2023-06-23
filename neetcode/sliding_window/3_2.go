package main

func characterReplacement(s string, k int) int {
	hash := map[byte]int{}
	l := 0
	res := 0
	maxF := 0
	for r, v := range s {
		hash[byte(v)] += 1
		if hash[byte(v)] > maxF {
			maxF = hash[byte(v)]
		}
		if (r-l+1)-maxF > k {
			hash[s[l]]--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
