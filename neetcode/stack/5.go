package main

import "fmt"

func main() {
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temp))
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	istack := make([]int, 0)
	ans := make([]int, len(temperatures))

	stack = append(stack, temperatures[0])
	istack = append(istack, 0)
	for i, v := range temperatures {
		if i == 0 {
			continue
		}
		if v <= stack[len(stack)-1] {
			stack = append(stack, v)
			istack = append(istack, i)
		} else {
			for len(stack)-1 >= 0 && v > stack[len(stack)-1] {
				itop := istack[len(istack)-1]
				ans[itop] = i - itop
				stack = stack[:len(stack)-1]
				istack = istack[:len(istack)-1]
			}
			stack = append(stack, v)
			istack = append(istack, i)
		}
	}
	return ans
}
