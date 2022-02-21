package main

import "fmt"

const size = 3

func determinant(m [size][size]int) int {
	detM00 := m[1][1]*m[2][2] - m[1][2]*m[2][1]
	detM01 := m[1][0]*m[2][2] - m[1][2]*m[2][0]
	detM03 := m[1][0]*m[2][1] - m[1][1]*m[2][0]
	det := m[0][0]*detM00 - m[0][1]*detM01 + m[0][2]*detM03
	return det
}

func main() {
	var matrix = [size][size]int{
		{1, 2, 3},
		{5, 8, 9},
		{6, 4, 3},
	}
	fmt.Println(determinant(matrix))
}
