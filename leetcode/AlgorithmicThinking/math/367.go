/*
	leetcode tag:367 title:有效的完全平方数
*/

package math

// method of math
func isPerfectSquare(num int) bool {
	r := num
	for r*r > num {
		r = (r + num/r) / 2
	}
	if r*r == num {
		return true
	}
	return false
}

// method of binary search
func isPerfectSquare2(num int) bool {
	l, r := 1, num
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid > num {
			r = mid - 1
			continue
		} else if mid*mid < num {
			l = mid + 1
			continue
		}
		return true
	}
	return false
}
