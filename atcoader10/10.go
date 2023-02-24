package main

import (
	"fmt"
	"math"
)

type Point struct {
	T int
	X int
	Y int
}

func main() {
	var N int
	fmt.Scan(&N)
	points := make([]Point, N+1)
	for i := 1; i < N+1; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		points[i].T, points[i].X, points[i].Y = t, x, y
	}
	points[0].T, points[0].X, points[0].Y = 0, 0, 0
	for i := range points {
		if i == 0 {
			continue
		}
		time := points[i].T - points[i-1].T
		dist := int(math.Abs(float64(points[i].X-points[i-1].X)) + math.Abs(float64(points[i].Y-points[i-1].Y)))

		if dist > time || dist%2 != time%2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
