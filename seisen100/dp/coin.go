package main

import "fmt"

func main() {
	INFTY := 50001
	var (
		n, m int
	)
	fmt.Scanf("%d %d", &n, &m)
	C := make([]int, m)
	for i := 0; i < m; i++ {
		var e int
		fmt.Scanf("%d", &e)
		C[i] = e
	}

	T := make([]int, n+1)
	T[0] = 0
	for i := 1; i < n+1; i++ {
		T[i] = INFTY
	}
	for _, cv := range C {
		for ti, tv := range T {
			if ti-cv >= 0 {
				T[ti] = min(tv, T[ti-cv]+1)
			}
		}
	}
	fmt.Println(T[n])
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
