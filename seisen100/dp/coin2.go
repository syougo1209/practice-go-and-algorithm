package main

import "fmt"

func main() {
	INFTY := 50001
	var (
		n, m int
	)
	fmt.Scanf("%d %d", &n, &m)
	coins := make([]int, m)
	for i := range coins {
		var c int
		fmt.Scanf("%d", &c)
		coins[i] = c
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := range dp[0] {
		dp[0][i] = INFTY
	}
	for i := range dp {
		dp[i][0] = 0
	}
	for ci, cv := range coins {
		for value := 0; value < n+1; value++ {
			if value-cv >= 0 {
				dp[ci+1][value] = min(dp[ci][value], dp[ci+1][value-cv]+1)
			} else {
				dp[ci+1][value] = dp[ci][value]
			}
		}
	}
	fmt.Println(dp[m][n])
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
