/*
	leetcode tag:540 title:有序数组中的单一元素
*/

package binarysearch

// method of binary search
func SingleNonDuplicate(nums []int) int {
	l, r, mid := 0, len(nums)-1, 0
	// 首先只有一个数是单条的，其余的都是成双的，那么数组长度肯定为奇数
	for l <= r {
		mid = l + (r-l)/2
		if mid-1 >= 0 && nums[mid] == nums[mid-1] {
			if (mid-1)%2 == 1 {
				r = mid - 2
			} else {
				l = mid + 1
			}
		} else if mid+1 < len(nums) && nums[mid] == nums[mid+1] {
			if (r-mid-1)%2 == 1 {
				l = mid + 2
			} else {
				r = mid - 1
			}
		} else {
			return nums[mid]
		}
	}

	return -1
}
