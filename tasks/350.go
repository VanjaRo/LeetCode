package main

import (
	"fmt"
)

func main() {
	arr1 := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}
	fmt.Print(intersect(arr1, arr2))
}

func intersect(nums1 []int, nums2 []int) []int {
	var ret []int
	numMap := make(map[int]int)
	for _, el := range nums1 {
		numMap[el]++
	}
	for _, el := range nums2 {
		if _, ok := numMap[el]; ok {
			numMap[el]--
			if numMap[el] == 0 {
				delete(numMap, el)
			}
			ret = append(ret, el)
		}

	}

	return ret
}
