/*
	题目描述：假设农场中成熟的母牛每年都会生 1 头小母牛，
			并且永远不会死。第一年有 1 只小母牛，从第二年
			开始，母牛开始生小母牛。每只小母牛 3 年之后成
			熟又可以生小母牛。给定整数 N，求 N 年后牛的数量。
*/

package dynamicprogramming

// method of dp
func ohter01(n int) int {
	if n < 3 {
		return 1
	}

	dp := make([]int, n)
	dp[0], dp[1], dp[2] = 1, 1, 1
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-3]
	}
	return dp[n-1]
}
