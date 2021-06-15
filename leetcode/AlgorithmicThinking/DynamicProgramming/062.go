/*
	leetcode tag:062 title:不同路径
*/

package dynamicprogramming

// method of dp
func uniquePaths(m int, n int) int {
	list := make([][]int, m)
	for i := 0; i < m; i++ {
		list[i] = make([]int, n)
		list[i][0] = 1
	}

	for i := 1; i < n; i++ {
		list[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			list[i][j] = list[i-1][j] + list[i][j-1]
		}
	}

	return list[m-1][n-1]
}
