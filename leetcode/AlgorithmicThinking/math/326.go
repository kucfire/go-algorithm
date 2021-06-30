/*
	leetcode tag:326 title:3的幂
*/

package math

// method of math
func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%3 == 0 {
		if isPowerOfThree(n / 3) {
			return true
		}
	}
	return false
}
