package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Print(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	l := len(nums)
	ret := make([]int, l)
	i, j, k := 0, l-1, l-1
	for i <= j {
		if modInt(nums[i]) < modInt(nums[j]) {
			ret[k] = nums[j] * nums[j]
			j--
		} else {
			ret[k] = nums[i] * nums[i]
			i++
		}
		k--
	}
	return ret
}

func modInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
