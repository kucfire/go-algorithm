/*
	leetcode tag:646 title:最长数对链
*/

package dynamicprogramming

import "sort"

// method of dp
func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	if n == 1 {
		return 1
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	dp := make([]int, n)
	dp[0] = 0
	r := 0

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if pairs[i][0] > pairs[j][1] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if r < dp[i] {
			r = dp[i]
		}
	}
	return r + 1
}
