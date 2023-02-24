package main

import (
	"fmt"
)

func main() {
	var A, B, C, X int
	fmt.Scan(&A, &B, &C, &X)

	var ans int

	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				total := i*500 + j*100 + 50*k
				if total == X {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
