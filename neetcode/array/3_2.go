package main

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		hashTable[v] = i
	}

	for i, v := range nums {
		if _, ok := hashTable[target-v]; !ok || hashTable[target-v] == i {
			continue
		} else {
			return []int{i, hashTable[target-v]}
		}
	}
	return []int{}
}
