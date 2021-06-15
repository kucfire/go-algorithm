/*
	leetcode tag:413 title:等差数列划分
*/

package dynamicprogramming

// method of dp
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	list := make([]int, len(nums))
	result := 0

	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			list[i] = list[i-1] + 1
			result += list[i]
		}
	}

	return result
}
