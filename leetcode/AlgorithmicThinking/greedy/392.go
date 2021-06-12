/*
	leetcode tag:392 title:判断子序列
*/

package greedy

// method of greedy
func isSubsequence(s string, t string) bool {
	pre := -1
loop1:
	for _, str := range s {
		if pre >= len(t) {
			return false
		}
		for i := pre + 1; i < len(t); i++ {
			if byte(str) == t[i] {
				pre = i
				continue loop1
			}
		}
		return false
	}
	return true
}
