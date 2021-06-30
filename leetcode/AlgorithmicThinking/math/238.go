/*
	leetcode tag:238 title:除自身以外数组的乘积
*/

package math

// method of math
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	tmp := 1
	result[0] = 1
	// 把左侧的乘积放入当前的位置
	for i := 1; i < len(nums); i++ {
		result[i] = tmp * nums[i-1]
		tmp = result[i]
	}

	tmp = 1
	// 将右侧的乘积合并入当前位置
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * tmp
		tmp = tmp * nums[i]
	}

	return result
}
