package main

import "fmt"

func leftKey(i int) int {
	return 2 * i
}
func rightKey(i int) int {
	return 2*i + 1
}
func parentKey(i int) int {
	return i / 2
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var e int
		fmt.Scan(&e)
		arr[i] = e
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("node %d: ", i)
		fmt.Printf("key %d: ", arr[i])
		if parentKey(i) > 0 {
			fmt.Printf("parent key %d: ", arr[parentKey(i)])
		}
		if leftKey(i) < len(arr) {
			fmt.Printf("left key %d: ", arr[leftKey(i)])
		}
		if rightKey(i) < len(arr) {
			fmt.Printf("right key %d: ", arr[rightKey(i)])
		}
		fmt.Printf("\n")
	}
	return
}
