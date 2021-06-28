/*
	leetcode tag:072 title:编辑距离
*/

package dynamicprogramming

// method of dp
func minDistance072(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	min := func(i ...int) int {
		res := i[0]
		for j := 1; j < len(i); j++ {
			if res > i[j] {
				res = i[j]
			}
		}
		return res
	}

	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
