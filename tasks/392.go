package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	left := 0
	for i := 0; i < len(t) && left < len(s); i++ {
		if t[i] == s[left] {
			left++
		}
	}
	return left == len(s)
}
