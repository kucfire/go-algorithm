/*
	leetcode tag:047 title:全排列 II
*/

package backtracking

import "sort"

// method of backtracking
func permuteUnique(nums []int) [][]int {
	results := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)

	isRepeat := make([]bool, n)

	dfs := func(resultStr []int) {}
	dfs = func(resultStr []int) {
		if len(resultStr) == len(nums) {
			results = append(results, resultStr)
			return
		}
		for i, num := range nums {
			if isRepeat[i] || (i > 0 && !isRepeat[i-1] && num == nums[i-1]) { // 重复判断
				continue
			}
			isRepeat[i] = true
			resultStr = append(resultStr, num)
			tmp := make([]int, len(resultStr))
			copy(tmp, resultStr)
			dfs(tmp)
			isRepeat[i] = false
			resultStr = resultStr[:len(resultStr)-1]
		}
	}

	dfs([]int{})

	return results
}
