package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		if v == string('+') || v == string('-') || v == string('*') || v == string('/') {
			calc := 0
			if v == string('+') {
				calc = stack[len(stack)-1] + stack[len(stack)-2]
			} else if v == string('-') {
				calc = stack[len(stack)-1] - stack[len(stack)-2]
			} else if v == string('*') {
				calc = stack[len(stack)-1] * stack[len(stack)-2]
			} else if v == string('/') {
				calc = stack[len(stack)-1] / stack[len(stack)-2]
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, calc)
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}
