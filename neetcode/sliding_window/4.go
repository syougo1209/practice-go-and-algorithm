package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	l := len(s1)

	for i := 0; i < len(s2); i++ {
		if i+l > len(s2) {
			break
		}
		substr := s2[i : i+l]
		if isPerm(s1, substr) {
			return true
		}
	}
	return false
}

func isPerm(s1 string, s2 string) bool {
	flg := map[rune]int{}
	for i := range s1 {
		flg[rune(s1[i])]++
		flg[rune(s2[i])]--
	}
	for i := range s1 {
		if flg[rune(s1[i])] != 0 {
			return false
		}
	}
	return true
}
