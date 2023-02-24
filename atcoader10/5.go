package main

import "fmt"

func sum(n int) int {
	var res int
	for 0 < n {
		res += n % 10
		n /= 10
	}
	return res
}

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)

	var ans int
	for i := 1; i <= N; i++ {
		s := sum(i)
		if A <= s && s <= B {
			ans += i
		}
	}
	fmt.Println(ans)
}
