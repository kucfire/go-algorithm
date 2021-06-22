/*
	leetcode tag:322 title:零钱兑换
*/

package dynamicprogramming

// method of dp
func coinChange(coins []int, amount int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	dp := make([]int, amount+1)
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if i == coin {
				dp[i] = 1
			} else if dp[i] == 0 && dp[i-coin] != 0 {
				dp[i] = dp[i-coin] + 1
			} else if dp[i] != 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}
