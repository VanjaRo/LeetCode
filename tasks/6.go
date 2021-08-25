package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {

	if len(s) <= numRows || numRows == 1 {
		return s
	}
	// min(diag) == 2
	// diag == numRows - 2
	maxStep := numRows + numRows - 2
	ret := make([]byte, len(s))

	currInd := 0
	currRow := 0
	for currRow < numRows {

		// first and last are symmetrical
		if currRow == 0 || currRow == numRows-1 {
			for i := currRow; i < len(s); i += maxStep {
				ret[currInd] = s[i]
				currInd++
			}
		} else {
			for i := currRow; i < len(s); i += maxStep {
				ret[currInd] = s[i]
				currInd++
				pairedElInd := i + maxStep - 2*currRow
				if pairedElInd < len(s) {
					ret[currInd] = s[pairedElInd]
					currInd++
				}
			}
		}
		currRow++
	}
	return string(ret)
}
