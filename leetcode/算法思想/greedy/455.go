/*
	leetcode tag:455 title:分发饼干
*/

package greedy

import "sort"

// method of greedy
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	result := 0
	indexG, indexS := 0, 0
	for indexG < len(g) && indexS < len(s) {
		if g[indexG] <= s[indexS] {
			result++
			indexG++
		}
		indexS++
	}
	return result
}
