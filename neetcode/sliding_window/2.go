package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	str := []byte(s)
	i := 0
	j := 0
	tmp := 0
	max := 0
	window := map[byte]bool{}
	n := len(str)
	for i < n && j < n {
		if window[str[j]] {
			max = maxi(max, tmp)
			delete(window, str[i])
			tmp--
			i++
		} else {
			window[str[j]] = true
			tmp++
			j++
		}
	}
	return maxi(max, tmp)
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}
