package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	r := bufio.NewReader(os.Stdin)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		A[i] = a
	}
	ans := 0
	var i int
	for {
		if A[i]%2 == 1 {
			break
		} else {
			A[i] = A[i] / 2
		}
		i++
		if i == n-1 {
			ans++
			i = 0
		}
	}
	fmt.Println(ans)
}
