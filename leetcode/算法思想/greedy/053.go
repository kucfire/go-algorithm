/*
	leetcode tag:053 title:最大子序和
*/

package greedy

import "math"

// method og greedy
func maxSubArray(nums []int) int {
	max := math.MinInt64
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1] // 当前所在位置的最大值
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
