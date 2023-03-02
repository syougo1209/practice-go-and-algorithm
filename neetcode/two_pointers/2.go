func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	for end > start {
		a := numbers[start] + numbers[end]
		if a == target {
			return []int{start + 1, end + 1}
		}
		if a > target {
			end--
		} else {
			start++
		}
	}
	return nil
}
