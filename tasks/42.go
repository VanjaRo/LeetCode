package main

import "fmt"

func main() {
	fmt.Print(trap([]int{4, 2, 0, 3, 2, 5}))
}

func trap(height []int) int {
	left := 0
	right := len(height) - 1

	ret := 0

	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				ret += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				ret += rightMax - height[right]
			}
			right--
		}
	}
	return ret
}
