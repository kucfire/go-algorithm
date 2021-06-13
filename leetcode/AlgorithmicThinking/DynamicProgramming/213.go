/*
	leetcode tag:213 title:打家劫舍 II
*/

package dynamicprogramming

// method of dp
func rob2(nums []int) int {
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

	// 从第一位开始
	firstOne, secondOne := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums)-1; i++ {
		firstOne, secondOne = secondOne, max(secondOne, firstOne+nums[i])
	}
	if len(nums) == 2 {
		return secondOne
	}

	// 从第二位开始
	firstTwo, secondTwo := nums[1], max(nums[1], nums[2])
	for i := 3; i < len(nums); i++ {
		firstTwo, secondTwo = secondTwo, max(secondTwo, firstTwo+nums[i])
	}

	return max(secondOne, secondTwo)
}
