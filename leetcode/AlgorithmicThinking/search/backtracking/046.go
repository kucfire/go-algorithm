/*
	leetcode tag:046 title:全排列
*/

package backtracking

// method of simple
func permute(nums []int) [][]int {
	result := make([][]int, 0)

	dfs := func(resultStr, lasts []int) {}
	dfs = func(resultStr, lasts []int) {
		if len(lasts) == 0 {
			result = append(result, resultStr)
			return
		}
		for i, num := range lasts {
			tmp := make([]int, len(lasts)-1) // ****
			copy(tmp[0:], lasts[:i])         // ****
			copy(tmp[i:], lasts[i+1:])       // ****
			dfs(append(resultStr, num), tmp)
		}
	}

	dfs([]int{}, nums)

	return result
}

// method of backtracking
func permuteBacktracking(nums []int) [][]int {
	result := make([][]int, 0)
	numsList := make(map[int]bool)

	dfs := func(resultStr []int) {}
	dfs = func(resultStr []int) {
		if len(resultStr) == len(nums) {
			result = append(result, resultStr)
			return
		}
		for _, num := range nums {
			if !numsList[num] {
				numsList[num] = true
				resultStr = append(resultStr, num)
				tmp := make([]int, len(resultStr))
				copy(tmp, resultStr)
				dfs(tmp)
				numsList[num] = false
				resultStr = resultStr[:len(resultStr)-1]
			}
		}
	}

	dfs([]int{})

	return result
}
