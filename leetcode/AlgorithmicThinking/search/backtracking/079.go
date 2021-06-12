/*
	leetcode tag:079 title:单词搜索
*/

package backtracking

// method of backtracking
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// 构建一个对应board的二维图，用来表示对应的位置是否路过
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	dps := func(i, j, l int) bool {
		return false
	}
	dps = func(i, j, l int) bool {
		// 只要有一个不符合，立马掐断
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[l] || vis[i][j] {
			return false
		}

		// 当探索的l与要被探索的单词长度一致，则返回成功
		if l == len(word)-1 {
			return true
		}

		vis[i][j] = true

		defer func() {
			vis[i][j] = false
		}()

		if dps(i+1, j, l+1) || dps(i-1, j, l+1) || dps(i, j+1, l+1) || dps(i, j-1, l+1) {
			return true
		}

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dps(i, j, 0) {
				return true
			}
		}
	}
	return false
}
