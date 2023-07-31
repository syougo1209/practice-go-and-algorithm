package main

import "fmt"

type Item struct {
	value  int
	weight int
}

func main() {
	var (
		n, W int
	)
	fmt.Scanf("%d %d", &n, &W)
	items := make([]Item, n)
	for i := 0; i < n; i++ {
		var v, w int
		fmt.Scanf("%d %d", &v, &w)
		items[i] = Item{value: v, weight: w}
	}
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := range dp[0] {
		dp[0][i] = 0
	}

	for itemIndex, item := range items {
		for w := 0; w < W+1; w++ {
			if w-item.weight >= 0 {
				dp[itemIndex+1][w] = max(dp[itemIndex][w], dp[itemIndex+1][w-item.weight]+item.value)
			} else {
				dp[itemIndex+1][w] = dp[itemIndex][w]
			}
		}
	}
	fmt.Println(dp[n][W])
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
