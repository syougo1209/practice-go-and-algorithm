package main

import (
	"fmt"
)

func main() {
	var N, Y int
	fmt.Scan(&N, &Y)
	clear := false
	for i := 0; i <= N; i++ { //10000
		if clear {
			break
		}
		for j := 0; j <= N-i; j++ { // 5000
			total := 10000*i + 5000*j + 1000*(N-i-j)
			if total == Y {
				fmt.Println(i, j, (N - i - j))
				clear = true
				break
			}
		}
	}
	if !clear {
		fmt.Println(-1, -1, -1)
	}
}
