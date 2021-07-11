package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{4, 5, 6}, 3)
	fmt.Print(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j = m - 1, n - 1
	for l := len(nums1) - 1; l >= 0; l-- {
		if j < 0 || i >= 0 && nums1[i] >= nums2[j] {
			nums1[l] = nums1[i]
			i--
		} else {
			nums1[l] = nums2[j]
			j--
		}
	}
}
