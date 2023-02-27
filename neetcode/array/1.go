package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 3, 5, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	nums_map := map[int]int{}
	for _, n := range nums {
		if _, ok := nums_map[n]; !ok {
			nums_map[n] = 1
		} else {
			return true
		}
	}
	return false
}
