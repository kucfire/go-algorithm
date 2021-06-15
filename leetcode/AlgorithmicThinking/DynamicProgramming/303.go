/*
	leetcode tag:303 title:区域和检索 - 数组不可变
*/

package dynamicprogramming

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left < 0 {
		left = 0
	}
	if right > len(this.nums) {
		right = len(this.nums)

	}

	result := 0

	for left <= right {
		if left == right {
			result += this.nums[left]
			break
		}
		result += this.nums[left] + this.nums[right]
		left++
		right--
	}
	return result
}
