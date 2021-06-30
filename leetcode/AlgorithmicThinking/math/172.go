/*
	leetcode tag:172 title:阶乘后的零
*/

package math

// method of math
func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n = n / 5
	}
	return res
}
