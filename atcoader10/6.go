package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := range A {
		var a int
		fmt.Scan(&a)
		A[i] = a
	}
	sort.Slice(A, func(i, j int) bool { return A[i] > A[j] })

	var alice, bob int
	for i, v := range A {
		if i%2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}

	fmt.Println(alice - bob)
}
