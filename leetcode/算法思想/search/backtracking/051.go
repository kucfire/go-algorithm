/*
	leetcode tag:051 title:N皇后
*/

package backtracking

// method of backtracking
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	list := make([]string, n)

	// 校验冲突
	inEare := func(i, j int) {
		// 横
		for m := 0; m < n; m++ {

		}
		// 纵
		for m := 0; m < n; m++ {

		}
		// 斜(四个方向)

	}

	dfs := func(i, j int) {}
	dfs = func(i, j int) {

	}
	dfs(0, 0)
	return result
}
