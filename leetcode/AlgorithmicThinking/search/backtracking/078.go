/*
	leetcode tag:078 title:子集
*/

package backtracking

// method of backtracking
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)

	dfs := func(start int) {}
	dfs = func(start int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			list = append(list, nums[i])
			dfs(i + 1)
			list = list[:len(list)-1]
		}
	}
	dfs(0)

	return result
}
