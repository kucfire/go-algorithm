/*
	leetcode tag:216 title:组合总和 III
*/

package backtracking

import "sort"

// method of backtracking
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)

	dfs := func(surplus, start int) {}
	dfs = func(surplus, start int) {
		if surplus == 0 && len(list) == k {
			tmp := make([]int, len(list))
			copy(tmp, list)
			sort.Ints(tmp)
			result = append(result, tmp)
			return
		}
		if surplus < 0 {
			return
		}
		for i := start; i < 10; i++ {
			list = append(list, i)
			dfs(surplus-i, start+1)
			list = list[:len(list)-1]
		}
	}

	dfs(n, 1)

	return result
}
