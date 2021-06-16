/*
	leetcode tag:279 title:完全平方数
*/

package dynamicprogramming

import "math"

// method of dp
func numSquares(n int) int {
	dp := make([]int, n+1)
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	for i := 1; i <= n; i++ {
		curMin := math.MaxInt32

		for j := 1; j*j <= i; j++ {
			curMin = min(curMin, dp[i-j*j])
		}

		dp[i] = curMin + 1
	}

	return dp[n]
}
