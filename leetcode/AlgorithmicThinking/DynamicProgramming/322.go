/*
	leetcode tag:322 title:零钱兑换
*/

package dynamicprogramming

import "math"

// method of dp
func coinChange(coins []int, amount int) int {
	if amount <= 0 && len(coins) <= 0 {
		return -1
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
