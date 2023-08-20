package main

import "fmt"

func main() {
	sets := [][]int{[]int{1, 1}, []int{1, 1}, []int{1, 1}}
	fmt.Println(pacificAtlantic(sets))
}

func pacificAtlantic(heights [][]int) [][]int {
	M := len(heights)
	N := len(heights[0])
	res := [][]int{}
	pac := make([][]bool, len(heights))
	for i := range pac {
		pac[i] = make([]bool, len(heights[0]))
	}
	at := make([][]bool, len(heights))
	for i := range at {
		at[i] = make([]bool, len(heights[0]))
	}

	var dfs func(int, int, [][]bool)
	dfs = func(y int, x int, visit [][]bool) {
		adjacents := [4][2]int{[2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}}
		visit[y][x] = true
		for _, v := range adjacents {
			ny := y + v[0]
			nx := x + v[1]
			if ny < 0 || nx < 0 || nx >= N || ny >= M || heights[ny][nx] < heights[y][x] || visit[ny][nx] {
				continue
			}
			dfs(ny, nx, visit)
		}
	}

	for i := 0; i < N; i++ {
		dfs(0, i, pac)
		dfs(M-1, i, at)
	}
	for i := 0; i < M; i++ {
		dfs(i, 0, pac)
		dfs(i, N-1, at)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if pac[i][j] && at[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
