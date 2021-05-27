/*
	leetcode tag:034 title:在排序数组中查找元素的第一个和最后一个位置
*/

package binarysearch

// method of binary search
func searchRange(nums []int, target int) []int {
	// 找起始点
	start, end := 0, len(nums)-1
	left := -1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			left = mid
			end = mid - 1
		}
	}
	// 找结束点
	start, end = 0, len(nums)-1
	right := -1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			right = mid
			start = mid + 1
		}
	}

	return []int{left, right}
}
