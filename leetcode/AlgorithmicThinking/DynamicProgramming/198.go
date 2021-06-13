/*
	leetcode tag:198 tttle:打家劫舍
*/

package dynamicprogramming

// method of dp
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	first, second := nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		first, second = second, max(second, first+nums[i])
	}

	return second
}
