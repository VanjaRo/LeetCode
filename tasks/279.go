package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(numSquares(4))
}

func numSquares(n int) int {
	count := make([]int, n+1)
	count[0] = 0
	count[1] = 1
	for i := 2; i <= n; i++ {
		count[i] = -1
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			count[i] = minIntModified(count[i], count[i-j*j]+1)
		}
	}
	return count[n]
}

func minIntModified(a, b int) int {
	if a == -1 || b < a {
		return b
	}
	return a
}
