package main

import (
	"fmt"
)

func main() {
	fmt.Print(lengthOfLongestSubstring("aab"))
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	mp := make([]int, 128)
	for i := range mp {
		mp[i] = -1
	}
	windowStart := -1

	for windowEnd, r := range s {
		if mp[r] > windowStart {
			windowStart = mp[r]
		}
		mp[r] = windowEnd
		if maxLen < windowEnd-windowStart {
			maxLen = windowEnd - windowStart
		}
	}
	return maxLen
}

// func lengthOfLongestSubstring(s string) int {
// 	counter := 0
// 	mp := make(map[rune]int)
// 	windowStart := 0

// 	for windowEnd, r := range s {
// 		mp[r]++
// 		for mp[r] > 1 {
// 			mp[rune(s[windowStart])]--
// 			if mp[rune(s[windowStart])] == 0 {
// 				delete(mp, rune(s[windowStart]))
// 			}
// 			windowStart++
// 		}
// 		counter = maxInt(counter, windowEnd-windowStart+1)
// 	}
// 	return counter
// }

// func maxInt(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
