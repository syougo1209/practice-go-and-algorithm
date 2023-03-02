package main

import "math"

func main() {

}

func maxProfit(prices []int) int {
	min := math.MaxInt32
	res := 0
	for _, price := range prices {
		if min < price {
			if res < price-min {
				res = price - min
			}
		} else {
			min = price
		}
	}
	return res
}
