package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	a := isValidColumn(board)
	b := isValidRow(board)
	c := isValidSubBoxs(board)

	if a && b && c {
		return true
	} else {
		return false
	}
}

func isValidRow(board [][]byte) bool { //yoko
	for j := 0; j < 9; j++ {
		flag := map[string]int{}
		for i := 0; i < 9; i++ {
			flag[string(board[i][j])] += 1
		}
		for i, v := range flag {
			if v > 1 && i != "." {
				return false
			}
		}
	}
	return true
}
func isValidColumn(board [][]byte) bool {
	for j := 0; j < 9; j++ {
		flag := map[string]int{}
		for i := 0; i < 9; i++ {
			flag[string(board[j][i])] += 1
		}
		for i, v := range flag {
			if v > 1 && i != "." {
				return false
			}
		}
	}
	return true
}

func isValidSubBoxs(board [][]byte) bool {
	starts := [][]int{[]int{0, 0}, []int{0, 3}, []int{0, 6}, []int{3, 0}, []int{3, 3}, []int{3, 6}, []int{6, 0}, []int{6, 3}, []int{6, 6}}
	for _, v := range starts {
		if !isValid(board, v) {
			return false
		}
	}
	return true
}

func isValid(board [][]byte, v []int) bool {
	flag := map[string]int{}
	for i := v[0]; i < v[0]+3; i++ {
		for j := v[1]; j < v[1]+3; j++ {
			flag[string(board[i][j])] += 1
		}
	}
	for i, v := range flag {
		if v > 1 && i != "." {
			fmt.Println(i)
			fmt.Println(v)
			return false
		}
	}
	return true
}
