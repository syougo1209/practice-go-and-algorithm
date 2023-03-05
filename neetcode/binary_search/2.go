package main

func searchMatrix(matrix [][]int, target int) bool {
	rowInd := searchRow(matrix, target)
	if rowInd == -1 {
		return false
	}
	if search(matrix[rowInd], target) != -1 {
		return true
	} else {
		return false
	}
}

func searchRow(matrix [][]int, target int) int {
	rowCount := len(matrix)
	rowLength := len(matrix[0])
	l := 0
	r := rowCount

	for r > l {
		mid := (l + r) / 2
		targetMatrix := matrix[mid]
		if targetMatrix[0] <= target && target <= targetMatrix[rowLength-1] {
			return mid
		} else if targetMatrix[0] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	ans := -1
	l := 0
	r := len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			ans = mid
			break
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}
