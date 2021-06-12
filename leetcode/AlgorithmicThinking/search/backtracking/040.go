/*
	leetcode tag:040 title:组合总和 II
*/

package backtracking

import "sort"

// method of backtracking
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	list := make([]int, 0)

	dfs := func(surplus, start int) {}
	dfs = func(surplus, start int) {
		if surplus == 0 {
			tmp := make([]int, len(list))
			copy(tmp, list)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			if surplus < candidates[i] {
				return
			} else {
				list = append(list, candidates[i])
				dfs(surplus-candidates[i], i+1)
				list = list[:len(list)-1]
			}
		}
	}

	dfs(target, 0)

	return result
}
