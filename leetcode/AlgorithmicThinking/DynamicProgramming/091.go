/*
	leetcode tag:091 title:解码方法
*/

package dynamicprogramming

// method of dp
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	n := len(s)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
