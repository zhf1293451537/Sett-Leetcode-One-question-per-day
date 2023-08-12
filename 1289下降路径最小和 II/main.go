package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		{-73, 61, 43, -48, -36},
		{3, 30, 27, 57, 10},
		{96, -76, 84, 59, -15},
		{5, -49, 76, 31, -7},
		{97, 91, 61, -46, 67},
	}
	res := minFallingPathSum(grid)
	fmt.Printf("非0偏移下降路径最小值为:%d\n", res)
}
func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	for i := 0; i < n; i++ {
		dp[0][i] = grid[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if k == j {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+grid[i][j])
			}
		}
	}
	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//回溯
// func minFallingPathSum(grid [][]int) int {
// 	n := len(grid)
// 	maxV := math.MaxInt
// 	dfs(grid, n, 0, 0, 0, &maxV)
// 	return maxV
// }

// func dfs(grid [][]int, n, x, y, res int, maxV *int) {
// 	if x == n {
// 		*maxV = min(*maxV, res)
// 		return
// 	}
// 	for i := 0; i < n; i++ {
// 		if i == y && x != 0 {
// 			continue
// 		}
// 		res += grid[x][i]
// 		// fmt.Println(x+1, i, res)
// 		dfs(grid, n, x+1, i, res, maxV)
// 		res -= grid[x][i]
// 	}
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
