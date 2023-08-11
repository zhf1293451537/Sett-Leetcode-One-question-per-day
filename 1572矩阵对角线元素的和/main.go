package main

import "fmt"

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := diagonalSum(mat)
	fmt.Printf("矩阵对角线的值为:%d\n", res)
}

func diagonalSum(mat [][]int) int {
	n := len(mat)
	res := 0
	x, y := 0, n-1
	for i := 0; i < n; i++ {
		res += mat[i][x] + mat[i][y]
		x++
		y--
	}
	if n%2 == 1 {
		res -= mat[n/2][n/2]
	}
	return res
}
