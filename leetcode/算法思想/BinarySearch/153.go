/*
	leetcode tag:153 title:寻找旋转排序数组中的最小值
*/

package binarysearch

// method of binary search
func findMin(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[end] {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return nums[start]
}
