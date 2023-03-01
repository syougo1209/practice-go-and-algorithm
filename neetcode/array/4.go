package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, v := range strs {
		sortedStr := sortString(v)
		anagrams[sortedStr] = append(anagrams[sortedStr], v)
	}

	l := len(anagrams)
	ans := make([][]string, l)
	index := 0
	for _, v := range anagrams {
		ans[index] = v
		index++
	}
	return ans
}

func sortString(str string) string {
	strArray := strings.Split(str, "")
	sort.Strings(strArray)
	return strings.Join(strArray, "")
}
