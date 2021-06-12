/*
	leetcode tag:417 title:太平洋大西洋水流问题
*/
package dfs

// method of dfs
func pacificAtlantic(heights [][]int) [][]int {
	result := make([][]int, 0)
	m, n, dir := len(heights), len(heights[0]), [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}
	// 找出能流向上左其中一边和下右其中一边
	dfs := func(heights, travelled [][]int, i, j int) {}
	dfs = func(heights, travelled [][]int, i, j int) {
		if travelled[i][j] == 1 {
			return
		}
		travelled[i][j] = 1
		for _, d := range dir {
			newI := i + d[0]
			newJ := j + d[1]
			if !inArea(newI, newJ) || heights[i][j] > heights[newI][newJ] || travelled[newI][newJ] == 1 {
				continue
			}
			dfs(heights, travelled, newI, newJ)
		}
	}

	pcf, alt := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		pcf[i] = make([]int, n)
		alt[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dfs(heights, pcf, i, 0)
		dfs(heights, alt, i, n-1)
	}

	for i := 0; i < n; i++ {
		dfs(heights, pcf, 0, i)
		dfs(heights, alt, m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pcf[i][j] == 1 && alt[i][j] == 1 {
				result = append(result, [][]int{{i, j}}...)
			}
		}
	}
	return result
}
