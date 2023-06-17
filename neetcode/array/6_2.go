package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	cond := map[string]bool{}
	for i := 0; i < 9; i++ { //row
		for j := 0; j < 9; j++ { //col
			num := board[i][j]
			if string(num) != "" {
				_, ok1 := cond[fmt.Sprintf("%d in %d row", num, i)]
				_, ok2 := cond[fmt.Sprintf("%d in %d col", num, j)]
				_, ok3 := cond[fmt.Sprintf("%d in %d - %d sub", num, (i/3), j/3)]
				if !ok1 || !ok2 || !ok3 {
					return false
				} else {
					cond[fmt.Sprintf("%d in %d row", num, i)] = true
					cond[fmt.Sprintf("%d in %d col", num, j)] = true
					cond[fmt.Sprintf("%d in %d - %d sub", num, (i/3), j/3)] = true
				}
			}
		}
	}
	return true
}
