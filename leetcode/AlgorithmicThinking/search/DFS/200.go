/*
	leetcode tag:200 title:岛屿数量
*/

package dfs

// method of BFS
func numIslands(grid [][]byte) int {
	result := 0
	dfs := func(grid [][]byte, i, j int) {}
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				result++
			}
		}
	}
	return result
}
