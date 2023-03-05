package main

import "fmt"

func main() {
	a := []int{40, 40, 41, 41}
	num := 0

	for ok := true; ok; ok = next_permutation(a) {
		num++
		fmt.Println(a)
	}
	fmt.Println(num)
}
func next_permutation(a []int) bool {
	return _permutation_pandita(a, func(x, y int) bool { return x < y })
}

func prev_permutation(a []int) bool {
	return _permutation_pandita(a, func(x, y int) bool { return y < x })
}

func _permutation_pandita(a []int, less func(x, y int) bool) bool {
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
