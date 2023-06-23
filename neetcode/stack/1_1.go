package main

func isValid(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := []byte{}
	for _, v := range []byte(s) {
		pair, ok := pairs[v]
		if !ok {
			stack = append(stack, v)
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != pair {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
