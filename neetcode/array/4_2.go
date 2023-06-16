package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}
	for _, str := range strs {
		sortedStr := sortString(str)
		if _, ok := hashMap[sortedStr]; !ok {
			hashMap[sortedStr] = []string{}
		}
		hashMap[sortedStr] = append(hashMap[sortedStr], str)
	}

	var res [][]string

	i := 0
	for _, anagrams := range hashMap {
		res = append(res, anagrams)
		i++
	}
	return res
}

func sortString(str string) string {
	strArray := strings.Split(str, "")
	sort.Strings(strArray)
	return strings.Join(strArray, "")
}
