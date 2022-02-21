package main

import "fmt"

const (
	size1 = 3
	size2 = 5
	size3 = 4
)

func matrixMultiplication(a [size1][size2]int, b [size2][size3]int) [size1][size3]int {
	var c [size1][size3]int
	for i := 0; i < size1*size3; i++ {
		row := i / size3
		col := i % size3
		for r := 0; r < size2; r++ {
			c[row][col] += a[row][r] * b[r][col]
		}
	}
	return c
}

func main() {
	var matrix1 = [size1][size2]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	var matrix2 = [size2][size3]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}

	fmt.Println(matrixMultiplication(matrix1, matrix2))
}
