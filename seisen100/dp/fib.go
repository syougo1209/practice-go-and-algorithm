package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fib(n, map[int]int{}))
}

func fib(n int, hashMap map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if _, ok := hashMap[n]; ok {
		return hashMap[n]
	}
	hashMap[n] = fib(n-1, hashMap) + fib(n-2, hashMap)
	return hashMap[n]
}
