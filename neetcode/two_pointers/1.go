package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "0P"
	fmt.Println(isPalindrome(str))
}
func isPalindrome(s string) bool {
	re1 := regexp.MustCompile("[^a-z0-9]")
	str := re1.ReplaceAllString(strings.ToLower(s), "")
	target := []byte(str)
	start := 0
	end := len(str) - 1
	for start < end {
		if target[start] != target[end] {
			return false
		}
		start++
		end--
	}
	return true
}
