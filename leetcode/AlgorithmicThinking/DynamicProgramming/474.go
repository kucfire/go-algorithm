/*
	leetcode tag:474 title: 一和零
*/

package dynamicprogramming

// method of dp
func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 || strs == nil {
		return 0
	}

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		ones, zeros := 0, 0
		for _, s := range str {
			if s == '0' {
				zeros++
			} else {
				ones++
			}
		}

		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}

	return dp[m][n]
}
