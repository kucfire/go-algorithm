/*
	leetcode tag:695 title:岛屿的最大面积
*/

package dfs

// method of bfs
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(grid, i, j))
			}
		}
	}
	return maxArea
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	area := 1
	grid[i][j] = 0 // 如果为1的经过这里需要变成0，代表着该地块已被记录
	area += dfs(grid, i+1, j)
	area += dfs(grid, i, j+1)
	area += dfs(grid, i-1, j)
	area += dfs(grid, i, j-1)
	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
