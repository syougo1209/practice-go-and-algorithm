package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}
func generateParenthesis(n int) []string {
	base := make([]byte, 0)
	for i := 0; i < n; i++ {
		base = append(base, []byte("(")...)
	}
	for i := 0; i < n; i++ {
		base = append(base, []byte(")")...)
	}
	res := make([]string, 0)
	for ok := true; ok; ok = next_permutation((base)) {
		fmt.Println(string(base))
		if isValid(string(base)) {
			res = append(res, string(base))
		}
	}

	return res
}

func isValid(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		pair, ok := pairs[char]
		if !ok {
			stack = append(stack, char)
			continue
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

func next_permutation(a []byte) bool {
	return _permutation_pandita(a, func(x, y byte) bool { return x < y })
}
func _permutation_pandita(a []byte, less func(x, y byte) bool) bool {
	i := len(a) - 2
	// Find the right most incresing elements a[i] and a[i+1]
	for 0 <= i && !less(a[i], a[i+1]) {
		i--
	}
	if i == -1 {
		return false
	}
	j := i + 1
	// Find the smallest element that is greater than a[i] in a[i+1:]
	for k := j + 1; k < len(a); k++ {
		// a[i] < a[k] && a[k] <= a[j]
		if less(a[i], a[k]) && !less(a[j], a[k]) {
			j = k
		}
	}
	a[i], a[j] = a[j], a[i]
	for p, q := i+1, len(a)-1; p < q; p, q = p+1, q-1 {
		a[p], a[q] = a[q], a[p]
	}
	return true
}
