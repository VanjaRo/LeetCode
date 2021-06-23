package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var ret []string
	start := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			end = nums[i]
		} else {
			if start != end {
				ret = append(ret, fmt.Sprintf("%d->%d", start, end))
			} else {
				ret = append(ret, fmt.Sprint(start))
			}
			start = nums[i]
			end = nums[i]
		}
	}
	if start != end {
		ret = append(ret, fmt.Sprintf("%d->%d", start, end))
	} else {
		ret = append(ret, fmt.Sprint(start))
	}
	return ret
}
