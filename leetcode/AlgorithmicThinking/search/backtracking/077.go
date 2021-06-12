/*
	leetcode tag:077 title:组合
*/

package backtracking

// method of backtracking
func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	list := make([]int, 0)

	dfs := func(i, j int) {}
	dfs = func(i, j int) {
		if j == 0 {
			tmp := make([]int, k)
			copy(tmp, list)
			result = append(result, tmp)
			return
		}

		for m := i; m <= n; m++ {
			list = append(list, m)
			dfs(m+1, j-1)
			list = list[:len(list)-1]
		}
	}

	dfs(1, k)

	return result
}
