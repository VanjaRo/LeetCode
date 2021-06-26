package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {
	var ret int
	start, end := 0, 0
	for prev := -1; start <= end && end < len(nums); end++ {
		if nums[end] != 1 {
			if prev == -1 {
				prev = end
			} else {
				ret = max(ret, end-start-1)
				start = prev + 1
				prev = end
			}
		}
	}
	ret = max(ret, end-start-1)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
