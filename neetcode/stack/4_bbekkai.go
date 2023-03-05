package main

import "strings"

func generateParenthesis(n int) []string {
	var ans []string
	stack := make([]string, 0)

	var backTrack func(int, int)
	backTrack = func(openN int, closeN int) {
		if openN == closeN && openN == n {
			ans = append(ans, strings.Join(stack, ""))
			return
		}
		if openN < n {
			stack = append(stack, "(")
			backTrack(openN+1, closeN)
			pop(&stack)
		}
		if closeN < openN {
			stack = append(stack, ")")
			backTrack(openN, closeN+1)
			pop(&stack)
		}
	}
	backTrack(0, 0)
	return ans
}

func pop(list *[]string) {
	*list = (*list)[:len(*list)-1]
}
