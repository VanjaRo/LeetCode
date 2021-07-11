package main

import "fmt"

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	for _, el := range matrix {
		fmt.Println(el)
	}
	fmt.Println("")
	rotate(matrix)

	for _, el := range matrix {
		fmt.Println(el)
	}
	fmt.Println("")

}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i], matrix[i][j] = matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i]

		}
	}
}
