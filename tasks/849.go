package main

import "fmt"

func main() {
	arr := []int{1, 0, 0, 0, 1, 0, 1}
	fmt.Println(maxDistToClosest(arr))
}

func maxDistToClosest(seats []int) int {
	max := 0
	lastInd := len(seats) - 1
	prev := -1
	for i := range seats {
		if seats[i] == 1 {
			if prev == -1 {
				max = maxInt(max, i)
			} else {
				max = maxInt(max, (i-prev)/2)
			}
			prev = i
		}
	}

	if prev != lastInd {
		max = maxInt(max, lastInd-prev)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
