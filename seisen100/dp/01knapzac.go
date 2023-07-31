package main

import "fmt"

type Item struct {
	value  int
	weight int
}

func main() {
	var (
		N, W int
	)
	fmt.Scanf("%d %d", &N, &W)
	items := make([]Item, N)
	for i := 0; i < N; i++ {
		var v, w int
		fmt.Scanf("%d %d", &v, &w)
		items[i] = Item{value: v, weight: w}
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 0; i < W+1; i++ {
		dp[0][i] = 0
	}

	for itemIndex, item := range items {
		for w := range dp[itemIndex+1] {
			if w-item.weight >= 0 {
				dp[itemIndex+1][w] = max(dp[itemIndex][w], dp[itemIndex][w-item.weight]+item.value)
			} else {
				dp[itemIndex+1][w] = dp[itemIndex][w]
			}
		}
	}
	fmt.Println(dp[N][W])
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
