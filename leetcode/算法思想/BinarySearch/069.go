/*
	leetcode tag:069 title: x 的平方根
*/

package binarysearch

// method of binary search
func mySqrt(x int) int {
	left, right, mid := 1, x, 0

	for left <= right {
		mid = left + (right-left)/2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if sqrt < mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
