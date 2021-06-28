/*
	leetcode tag:650 title:只有两个键的键盘
*/

package dynamicprogramming

import "math"

func minSteps(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 2

	for i := 3; i <= n; i++ {
		minV := math.MaxInt32
		for j := i / 2; j > 2; j-- {
			if i%j == 0 {
				minV = dp[j] + i/j
				break
			}
		}
		dp[i] = minV
	}

	return dp[n]
}
