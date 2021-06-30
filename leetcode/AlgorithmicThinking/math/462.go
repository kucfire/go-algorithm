/*
	leetcode tag:462 title:最少移动次数使数组元素相等 II
*/

package math

import (
	"sort"
)

// method of math
func MinMoves2(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	mid := n / 2
	sum := 0
	fix := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	for i, num := range nums {
		if i == mid {
			continue
		}
		sum += fix(nums[mid] - num)
	}

	return sum
}
