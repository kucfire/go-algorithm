/*
	leetcode tag:204 title:计数质数
*/

package math

// method of math
func CountPrimes(n int) int {
	count := 0
	list := make([]bool, n)
	for i := 2; i < n; i++ {
		if list[i] {
			continue
		}
		for j := i; j < n; j += i {
			list[j] = true
		}
		count++
	}
	return count
}
