package main

import "strconv"

func evalRPN(tokens []string) int {
	st := make([]int, 0)
	push := func(v int) {
		st = append(st, v)
	}
	pop := func() int {
		n := len(st) - 1
		res := st[n]
		st = st[:n]
		return res
	}

	for _, v := range tokens {
		if v == "+" || v == "/" || v == "-" || v == "*" {
			a := pop()
			b := pop()
			if v == "+" {
				push(b + a)
			} else if v == "/" {
				push(b / a)
			} else if v == "-" {
				push(b - a)
			} else if v == "*" {
				push(b * a)
			}
		} else {
			i, _ := strconv.Atoi(v)
			push(i)
		}
	}
	return st[len(st)] - 1
}
