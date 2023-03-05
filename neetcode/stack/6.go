package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i, v := range position {
		cars[i] = Car{v, speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool { return cars[i].position > cars[j].position })
	fmt.Println(cars)
	stack := []Car{}
	stack = append(stack, cars[0])
	for i, v := range cars {
		if i == 0 {
			continue
		}
		top := stack[len(stack)-1]
		topt := float32(target-top.position) / float32(top.speed)
		vt := float32(target-v.position) / float32(v.speed)
		if topt < vt {
			stack = append(stack, v)
		}
	}
	return len(stack)
}

type Car struct {
	position int
	speed    int
}
