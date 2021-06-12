/*
	leetcode tag:051 title:N皇后
*/

package backtracking

// method of backtracking
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)

	// 构建一个棋盘
	boards := make([][]bool, n)
	for i := range boards {
		boards[i] = make([]bool, n)
	}

	// 校验冲突
	inEare := func(i, j int) bool {
		// 纵横校验
		for m := 0; m < n; m++ {
			// 横向
			if m != i && boards[m][j] {
				return true
			}
			// 纵向
			if m != j && boards[i][m] {
				return true
			}
		}
		// 斜校验(四个方向)
		for m := 1; m <= n; m++ {
			// 超范围退出条件
			if i+m >= n && i-m < 0 && j+m >= n && j-m < 0 {
				break
			}
			// 左上
			if i-m >= 0 && j+m < n {
				if boards[i-m][j+m] {
					return true
				}
			}
			// 左下
			if i+m < n && j+m < n {
				if boards[i+m][j+m] {
					return true
				}
			}
			// 右上
			if i-m >= 0 && j-m >= 0 {
				if boards[i-m][j-m] {
					return true
				}
			}
			// 右下
			if i+m < n && j-m >= 0 {
				if boards[i+m][j-m] {
					return true
				}
			}
		}

		return false
	}

	dfs := func(i, j, QNum int) {}
	dfs = func(i, j, QNum int) {
		if QNum == n {
			resultStr := make([]string, 0)
			for _, board := range boards {
				resultStrStr := ""
				for _, Q := range board {
					if Q {
						resultStrStr += "Q"
					} else {
						resultStrStr += "."
					}
				}
				resultStr = append(resultStr, resultStrStr)
			}
			result = append(result, resultStr)
			return
		}

		// i,j大小校验
		if j >= n {
			i++
			if i >= n {
				return
			}
			j = 0
		}

		if !inEare(i, j) {
			QNum++
			boards[i][j] = true
			dfs(i, j+1, QNum)

			// 回溯
			QNum--
			boards[i][j] = false
		}
		dfs(i, j+1, QNum)
	}
	dfs(0, 0, 0)
	return result
}
