/*
	leetcode tag:416 title:分割等和子集
*/

package dynamicprogramming

// 题意是要找出和相等的两个子集
// 那么和相等的意思就是数组的总是两个子集和的总和

// method of dp
func canPartition(nums []int) bool {
	// 需要先计算数组内总和是否为偶数，不为偶数的话无法分割出等和的子集
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	w := sum / 2
	dp := make([]bool, w+1)
	dp[0] = true

	for _, num := range nums {
		for i := w; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[w]
}
