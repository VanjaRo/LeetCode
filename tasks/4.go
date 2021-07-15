package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 3}
	arr_2 := []int{2}
	fmt.Println(findMedianSortedArrays(arr1, arr_2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	medInd := len(nums1) / 2
	if len(nums1) == medInd*2 {
		return float64(nums1[medInd]+nums1[medInd-1]) / 2
	} else {
		return float64(nums1[medInd])
	}
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	merged := []int{}
// 	left1 := 0
// 	left2 := 0
// 	for left1 < len(nums1) {
// 		if left2 < len(nums2) {
// 			if nums1[left1] < nums2[left2] {
// 				merged = append(merged, nums1[left1])
// 				left1++
// 			} else {
// 				merged = append(merged, nums2[left2])
// 				left2++
// 			}
// 		} else {
// 			break
// 		}
// 	}
// 	if left2 < len(nums2) {
// 		merged = append(merged, nums2[left2:]...)
// 	}
// 	if left1 < len(nums1) {
// 		merged = append(merged, nums1[left1:]...)
// 	}

// 	medInd := len(merged) / 2
// 	if len(merged) == medInd*2 {
// 		return float64(merged[medInd]+merged[medInd-1]) / 2
// 	} else {
// 		return float64(merged[medInd])
// 	}
// }
