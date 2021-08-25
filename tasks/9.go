package main

import "fmt"

func main() {

}

func isPalindrome(x int) bool {
	sX := fmt.Sprintf("%d", x)
	for i := 0; i < len(sX)/2; i++ {
		if sX[i] != sX[len(sX)-1-i] {
			return false
		}
	}
	return true
}
