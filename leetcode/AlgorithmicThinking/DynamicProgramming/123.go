/*
	leetcode tag:123 title:买卖股票的最佳时机 III
*/

package dynamicprogramming

// method of dp
func maxProfitIII(prices []int) int {
	n := len(prices)

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0

	for i := 1; i < n; i++ {
		newbuy1 := max(buy1, -prices[i])
		newsell1 := max(sell1, buy1+prices[i])
		newbuy2 := max(buy2, sell1-prices[i])
		newsell2 := max(sell2, buy2+prices[i])
		buy1, sell1, buy2, sell2 = newbuy1, newsell1, newbuy2, newsell2
	}

	return sell2
}
