/*
	leetcode tag:139 title:单词拆分
*/

package dynamicprogramming

// method of dp
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	wordDictSet := make(map[string]bool)
	for _, word := range wordDict {
		wordDictSet[word] = true
	}

	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
