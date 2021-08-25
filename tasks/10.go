package main

import (
	"fmt"
)

func main() {
	fmt.Print(isMatch("absabs", "abs*"))
}

// abab	bas
// ab*	b.*

func isMatch(s string, p string) bool {
	m := make(map[[2]int]bool)
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if _, ok := m[[2]int{i, j}]; !ok {
			var ans bool
			if j == len(p) {
				ans = i == len(s)
			} else {
				firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')
				if j+1 < len(p) && p[j+1] == '*' {
					ans = dp(i, j+2) || firstMatch && dp(i+1, j)
				} else {
					ans = firstMatch && dp(i+1, j+1)
				}
			}
			m[[2]int{i, j}] = ans
		}
		return m[[2]int{i, j}]
	}
	return dp(0, 0)
}
