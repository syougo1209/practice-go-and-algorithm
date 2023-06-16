package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}
	for _, v := range strs {
		sortedString := sortString(s)
		if _, ok := hashMap[sortedString]; !ok {
			hashMap[sortedString] = []string{}
		}
		hashMap[sortedString] = append(hashMap[sortedString], v)
	}

	res := make([][]string, len(hashMap))
	i := 0
	for _, v := range hashMap {
		res[i] = v
		i++
	}
	return res
}

func sortString(str string) string {
	strarr := strings.Split(str, "")
	sort.Strings(strarr)
	return strings.Join(strarr, "")
}
