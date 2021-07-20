package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Print(searchRange(nums, 0))
}

func searchRange(nums []int, target int) []int {
	ind := len(nums) / 2
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if target < nums[ind] {
		for ; ind >= 0 && target != nums[ind]; ind-- {
			if target > nums[ind] {
				return []int{-1, -1}
			}
		}
		if ind < 0 {
			return []int{-1, -1}
		}
	}
	if target > nums[ind] {
		for ; ind < len(nums) && target != nums[ind]; ind++ {
			if target < nums[ind] {
				return []int{-1, -1}
			}
		}
		if ind >= len(nums) {
			return []int{-1, -1}
		}
	}
	return expandFromMiddle(&nums, target, ind)
}

func expandFromMiddle(nums *[]int, target, ind int) []int {
	right := ind
	left := ind
	for ; right < len(*nums); right++ {
		if (*nums)[right] != target {
			break
		}
	}

	for ; left >= 0; left-- {
		if (*nums)[left] != target {
			break
		}
	}
	left++
	right--
	return []int{left, right}

}
