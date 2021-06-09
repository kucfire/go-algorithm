/*
	leetcode tag:037 title:解数独
*/

package backtracking

func solveSudoku(board [][]byte) {
	// 重复检验
	isRepeat := func(i, j int, num byte) bool {
		// block的起点
		blocki := (i / 3) * 3
		blockj := (j / 3) * 3

		// 判断block范围内有没有重复的数字
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if m == i && n == j {
					continue
				}
				if num == board[blocki+m][blockj+n] {
					return true
				}
			}
		}
		// 判断纵横方向有没有重复的
		for m := 0; m < 9; m++ {
			// 横向
			if m != i && num == board[m][j] {
				return true
			}
			// 纵向
			if m != j && num == board[i][m] {
				return true
			}
		}
		return false
	}

	dfs := func(i, j int) bool { return true }
	dfs = func(i, j int) bool {
		if j == 9 {
			i++
			j = 0
			if i == 9 {
				return true
			}
		}

		return true
	}
	dfs(0, 0)
}
