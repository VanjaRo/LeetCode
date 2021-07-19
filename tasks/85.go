package main

import "fmt"

func main() {
	// matrix := [][]byte{
	// 	{'1', '1', '1', '1', '1', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1', '1', '1', '0'},
	// 	{'1', '1', '1', '1', '1', '1', '1', '0'},
	// 	{'1', '1', '1', '1', '1', '0', '0', '0'},
	// 	{'0', '1', '1', '1', '1', '0', '0', '0'}}

	// matrix2 := [][]byte{
	// 	{'1', '0', '0', '0', '1'},
	// 	{'1', '1', '0', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// }

	matr3 := [][]byte{
		{'1', '0', '1', '1', '1'},
		{'0', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '1'},
		{'1', '1', '0', '1', '1'},
		{'0', '1', '1', '1', '1'},
	}
	fmt.Print(maximalRectangle(matr3))
}

func maximalRectangle(matrix [][]byte) int {
	area := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				newArea := getArea(matrix, i, j)
				if area < newArea {
					area = newArea
				}
			}
		}
	}
	return area
}

func getArea(matrix [][]byte, x, y int) int {
	maxArea := 0
	minWidth := 0
	width := 0
	length := 0
	for i := x; i < len(matrix); i++ {
		if matrix[i][y] == '0' {
			break
		}
		for j := y; j < len(matrix[0]); j++ {
			if minWidth != 0 && j-y >= minWidth || matrix[i][j] == '0' {
				break
			}

			width++
		}

		length++
		if minWidth == 0 || width < minWidth {
			minWidth = width
		}
		if width*length > maxArea {
			maxArea = width * length
		}
		width = 0
	}
	return maxArea
}
