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
			chars[pivot] = chars[start]
			pivot++
			if start != end {
				for _, x := range fmt.Sprint(end - start + 1) {
					chars[pivot] = byte(x)
					pivot++
				}
			}
			start = i
		}
		end = i
	}
	chars[pivot] = chars[start]
	pivot++
	if start != end {
		for _, x := range fmt.Sprint(end - start + 1) {
			chars[pivot] = byte(x)
			pivot++
		}
	}
	start = i
	return pivot
}
