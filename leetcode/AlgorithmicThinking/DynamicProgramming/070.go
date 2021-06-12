/*
	leetcode tag:070 title:爬楼梯
*/

package dynamicprogramming

// method of fabonacci
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	result1 := 2
	result2 := 1
	for i := 2; i < n; i++ {
		cur := result1 + result2
		result2 = result1
		result1 = cur
	}
	return result1
}
