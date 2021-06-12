/*
	leetcode tag:278 title:第一个错误的版本
*/

package binarysearch

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// method of binary search
func firstBadVersion(n int) int {
	head, tail := 1, n
	for head <= tail {
		mid := head + (tail-head)/2
		if isBadVersion(mid) {
			if !isBadVersion(mid - 1) {
				return mid
			}
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}
	return -1
}

func isBadVersion(mid int) bool {
	return true
}
