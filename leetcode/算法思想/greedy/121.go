/*
	leetcode tag:121 title:买卖股票的最佳时机
*/

package greedy

import "math"

// method of greedy
func maxProfit(prices []int) int {
	result := 0
	minPrice := math.MaxInt64

	for _, price := range prices {
		if result < price-minPrice {
			result = price - minPrice
		} else if price < minPrice {
			minPrice = price
		}
	}

	return result
}
