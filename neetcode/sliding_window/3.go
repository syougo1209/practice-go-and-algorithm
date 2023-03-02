package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	flq := map[rune]int{}

	l := 0
	mfreq := 0
	ans := 0
	for r, v := range s {
		flq[v] += 1
		mfreq = maxi(mfreq, flq[v])
		if r-l+1-mfreq > k {
			flq[rune(s[l])]--
			l++
		} else {
			ans = maxi(ans, r-l+1)
		}
	}
	return ans
}

func maxi(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
