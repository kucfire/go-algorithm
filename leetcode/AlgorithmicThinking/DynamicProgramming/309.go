/*
	leetcode tag:309 title:最佳买卖股票时机含冷冻期
*/

package dynamicprogramming

// method of dp
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		newf0 := max(f0, f2-prices[i]) // 买入的后的收益
		newf1 := f0 + prices[i]        // 卖出后的收益
		newf2 := max(f1, f2)           // 比较刚卖出和卖出第二天的收益大小，取最大的
		f0, f1, f2 = newf0, newf1, newf2
	}

	return max(f1, f2)
}
