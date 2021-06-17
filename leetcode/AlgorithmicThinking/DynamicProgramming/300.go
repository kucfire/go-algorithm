/*
	leetcode tag:300 title:最长递增子序列
*/

package dynamicprogramming

// method of dp
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}
