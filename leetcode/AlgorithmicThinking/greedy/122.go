/*
	leetcode tag:122 title:买卖股票的最佳时机 II
*/

package greedy

// method of greedy
func maxProfit2(prices []int) int {
	result := 0

	for i := range prices[:len(prices)-1] {
		if prices[i] < prices[i+1] {
			result += prices[i+1] - prices[i]
		}
	}

	return result
}
