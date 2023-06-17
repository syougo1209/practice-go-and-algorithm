package main

import "math"

func maxProfit(prices []int) int {
	min := math.MaxUint32
	res := 0
	for _, price := range prices {
		if price < min {
			min = price
		} else {
			if price-min > res {
				res = price - min
			}
		}
	}
	return res
}
