package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dp := map[int]int{}
	fmt.Println(fib(n, dp))
}

func fib(n int, dp map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if _, ok := dp[n]; ok {
		return dp[n]
	}
	dp[n] = fib(n-1, dp) + fib(n-2, dp)
	return dp[n]
}
