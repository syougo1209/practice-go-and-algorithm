package main

func lengthOfLongestSubstring(s string) int {
	res := 0
	hash := map[byte]bool{}
	l := 0

	for r, _ := range s {
		for hash[s[r]] {
			delete(hash, s[l])
			l++
		}
		hash[s[r]] = true
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
