package main

import "fmt"

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	start := 0
	end := 0
	pivot := 0
	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[i-1] {
			if start != end {
				chars[pivot] = chars[start]
				pivot++
				for _, x := range fmt.Sprint(end - start + 1) {
					chars[pivot] = byte(x)
					pivot++
				}
			} else {
				chars[pivot] = chars[start]
				pivot++
			}
			start = i
		}
		end = i
	}
	if start != end {
		chars[pivot] = chars[start]
		pivot++
		for _, x := range fmt.Sprint(end - start + 1) {
			chars[pivot] = byte(x)
			pivot++
		}
	} else {
		chars[pivot] = chars[start]
		pivot++
	}
	return pivot
}
