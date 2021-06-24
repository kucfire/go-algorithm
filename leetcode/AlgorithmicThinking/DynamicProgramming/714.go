/*
	leetcode tag:714 title:买卖股票的最佳时机含手续费
*/

package dynamicprogramming

// method of dp
func maxProfitCommission(prices []int, fee int) int {
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

	f0, f1 := -prices[0], 0
	for i := 1; i < n; i++ {
		newf0 := max(f1-prices[i], f0)
		newf1 := max(f0+prices[i]-fee, f1)
		f0, f1 = newf0, newf1
	}

	return f1
}
