package main

func main() {
	nums := []int{2, 7, 11, 15}
	twoSum(nums, 9)
}
func twoSum(nums []int, target int) []int {
	num_index := map[int]int{}
	for _, v := range nums {
		num_index[v] += 1
	}
	var sum1, sum2 int
	for _, v := range nums {
		if _, ok := num_index[target-v]; ok {
			if target == 2*v {
				if num_index[v] <= 1 {
					continue
				}
			}
			sum1 = v
			sum2 = target - v
			break
		}
	}
	var ans1 int
	for i, v := range nums {
		if v == sum1 {
			ans1 = i
			break
		}
	}
	var ans2 int
	for i, v := range nums {
		if v == sum2 && i != ans1 {
			ans2 = i
		}
	}

	return []int{ans1, ans2}
}
