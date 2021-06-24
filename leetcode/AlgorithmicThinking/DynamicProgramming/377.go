/*
	leetcode tag:377 title:组合总和
*/

package dynamicprogramming

// method of dp
func combinationSum4(nums []int, target int) int {
	if target < 1 && len(nums) < 1 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
