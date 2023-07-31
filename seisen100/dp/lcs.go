package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d %d", &n)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		var x, y string
		fmt.Scanf("%s", &x)
		fmt.Scanf("%s", &y)
		res[i] = lcs(x, y)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func lcs(x string, y string) int {
	dp := make([][]int, len(y)+1)
	for i := range dp {
		dp[i] = make([]int, len(x)+1)
	}
	for i := 0; i < len(x)+1; i++ {
		dp[0][i] = 0
	}
	for i := 0; i < len(y)+1; i++ {
		dp[i][0] = 0
	}

	for iy, vy := range y {
		for ix, vx := range x {
			if vy == vx {
				dp[iy+1][ix+1] = dp[iy][ix] + 1
			} else {
				dp[iy+1][ix+1] = max(dp[iy][ix+1], dp[iy+1][ix])
			}
		}
	}
	return dp[len(y)][len(x)]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
