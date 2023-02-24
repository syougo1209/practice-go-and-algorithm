package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	D := make([]int, N)
	for i := range D {
		var d int
		fmt.Scan(&d)
		D[i] = d
	}
	diexist := make([]int, 100)
	for _, v := range D {
		diexist[v-1]++
	}
	var ans int
	for i := range diexist {
		if diexist[i] > 0 {
			ans++
		}
	}
	fmt.Println(ans)
}
