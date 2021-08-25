package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {

	strX := fmt.Sprintf("%d", x)
	// last zero
	if strX[len(strX)-1] == '0' {
		strX = strX[:len(strX)-1]
	}

	lenAfter := len(strX)
	if lenAfter > 0 {
		sign := 1
		// sign
		if strX[0] == '-' {
			sign = -1
			lenAfter--
		}

		ret := make([]byte, lenAfter)
		for i := 0; i < len(ret); i++ {
			ret[i] = strX[len(strX)-1-i]
		}
		val, _ := strconv.Atoi(string(ret))
		if val > 1<<31-1 {
			val = 0
		}
		return val * sign
	}

	return 0
}
