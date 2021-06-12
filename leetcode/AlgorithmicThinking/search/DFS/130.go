/*
	leetcode tag:130 title:被围绕的区域
*/

package dfs

// method of dfs
func solve(board [][]byte) {
	dfs := func(board [][]byte, i, j int) {}
	dfs = func(board [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 'O' {
			return
		}

		board[i][j] = 'A'
		dfs(board, i+1, j)
		dfs(board, i-1, j)
		dfs(board, i, j+1)
		dfs(board, i, j-1)
	}

	// 遍历两次
	// 第一次，找出边界及其相连的'O'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1) && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	// 第二次遍历，将所有'A'修改成'O'，将所有'O'修改成'X'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'A':
				board[i][j] = 'O'
			}
		}
	}
}
