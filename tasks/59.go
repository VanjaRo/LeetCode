package main

import "fmt"

func main() {
	fmt.Print(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	x := 0
	y := 0
	counter := 1
	for counter != n*n+1 {
		for x < n {
			if matrix[y][x] != 0 {
				break
			}
			matrix[y][x] = counter
			counter++
			x++
		}
		x--
		y++
		for y < n {
			if matrix[y][x] != 0 {
				break
			}
			matrix[y][x] = counter
			counter++
			y++
		}
		y--
		x--
		for x >= 0 {
			if matrix[y][x] != 0 {
				break
			}
			matrix[y][x] = counter
			counter++
			x--
		}
		x++
		y--
		for y >= 0 {
			if matrix[y][x] != 0 {
				break
			}
			matrix[y][x] = counter
			counter++
			y--
		}
		y++
		x++
	}
	return matrix
}
