/*
	leetcode tag:343 title:整数拆分
*/

package dynamicprogramming

// method of dp
func integerBreak(n int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}

	return dp[n]
}
