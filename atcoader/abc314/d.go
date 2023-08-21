package main

import (
	"fmt"
	"strings"
)

type Ope struct {
	t int
	x int
	c string
}

func main() {
	var N, Q int
	var S string

	fmt.Scan(&N)
	fmt.Scanf("%s", &S)
	fmt.Scan(&Q)
	ope := []Ope{}
	for i := 0; i < Q; i++ {
		var t, x int
		var c string
		fmt.Scanf("%d %d %s", &t, &x, &c)
		ope = append(ope, Ope{t: t, x: x - 1, c: c})
	}
	changeIndex := -1
	changeWay := 0
	for i := Q - 1; i >= 0; i-- {
		if ope[i].t != 1 && changeWay == 0 {
			changeWay = ope[i].t
			changeIndex = i
			break
		}
	}
	for i, v := range ope {
		if changeIndex == i {
			if changeWay == 2 {
				S = strings.ToLower(S)
			} else if changeWay == 3 {
				S = strings.ToUpper(S)
			}
			continue
		}
		if v.t == 1 {
			S = S[:v.x] + v.c + S[v.x+1:]
		}
	}
	fmt.Println(S)
}
