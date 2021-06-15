/*
	leetcode tag:064 title:最小路径和
*/

package dynamicprogramming

import "math"

// method of dp
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			up, left := math.MaxInt32, math.MaxInt32
			if i-1 >= 0 {
				left = grid[i-1][j] + grid[i][j]
			}
			if j-1 >= 0 {
				up = grid[i][j-1] + grid[i][j]
			}
			grid[i][j] = min(up, left)
		}
	}
	return grid[n-1][m-1]
}
