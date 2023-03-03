package main

import "fmt"

func main() {
	fmt.Println(isValid("(]"))
}
func isValid(s string) bool {
	st := make([]rune, 0)
	push := func(v rune) {
		st = append(st, v)
	}
	pop := func() rune {
		n := len(st) - 1
		res := st[n]
		st = st[:n]
		return res
	}
	empty := func() bool {
		return len(st) == 0
	}
	for _, v := range s {
		if empty() {
			push(v)
		} else if isTwin(st[len(st)-1], v) {
			pop()
		} else {
			push(v)
		}
	}
	return empty()
}

func isTwin(s1 rune, s2 rune) bool {
	if (string(s2) == ")" && string(s1) == "(") || (string(s2) == "}" && string(s1) == "{") || (string(s2) == "]" && string(s1) == "[") {
		return true
	} else {
		return false
	}
}
