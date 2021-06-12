/*
	leetcode tag:090 title:子集 II
*/

package backtracking

import "sort"

// method of backtracking
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	list := make([]int, 0)

	dfs := func(start int) {}
	dfs = func(start int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			list = append(list, nums[i])
			dfs(i + 1)
			list = list[:len(list)-1]
		}
	}
	dfs(0)

	return result
}
