package main

func lengthOfLongestSubstring(s string) int {
	sb := []byte(s)
	res := 0
	cnt := 1
	i := 0
	for i < len(sb) {
		cnt = 1
		hash := map[byte]bool{}
		hash[sb[i]] = true
		j := i + 1
		for j < len(sb) {
			if _, ok := hash[sb[j]]; ok {
				break
			} else {
				hash[sb[j]] = true
				j++
				cnt++
			}
		}
		if cnt > res {
			res = cnt
		}
		i++
	}
	return res
}
