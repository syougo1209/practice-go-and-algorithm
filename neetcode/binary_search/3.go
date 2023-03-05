package main

func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := 1000000000 + 1
	var mid int
	ans := 1
	for l < r {
		mid = (l + r) / 2
		if isEat(piles, h, mid) { //早すぎる
			ans = mid
			r = mid
		} else { //遅すぎる
			l = mid + 1
		}
	}
	return ans
}

func isEat(piles []int, h int, s int) bool {
	t := 0
	for _, v := range piles {
		if v%s == 0 {
			t += v / s
		} else {
			t += v/s + 1
		}
	}
	return t <= h
}
