/*
	leetcode tag:628 title:三个数的最大乘积
*/

package math

import "sort"

// method of math
func maximumProduct(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	max := func(i int, j int) int {
		if i > j {
			return i
		}
		return j
	}
	return max(nums[n-1]*nums[n-2]*nums[n-3], nums[0]*nums[1]*nums[n-1])
}
