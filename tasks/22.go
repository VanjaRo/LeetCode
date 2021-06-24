package main

import "fmt"

func main() {
	fmt.Print(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var ret []string
	backTrack(&ret, "", n, n, n*2)
	return ret

}

func backTrack(output *[]string, current string, opened, closed, max int) {
	if len(current) == max {
		*output = append(*output, current)
	}
	if opened > 0 {
		backTrack(output, current+"(", opened-1, closed, max)

	}
	if closed > opened {
		backTrack(output, current+")", opened, closed-1, max)
	}
}
