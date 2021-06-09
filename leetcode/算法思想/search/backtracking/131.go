/*
	leetcode tag:131 title:分割回文串
*/

package backtracking

// method of backtracking
func partition(s string) [][]string {
	result := make([][]string, 0)
	list := make([]string, 0)
	// sArray := strings.Split(s, "")

	isPartition := func(str string) bool {
		for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}

	dfs := func(start int) {}
	dfs = func(start int) {
		if start == len(s) {
			tmp := make([]string, len(list))
			copy(tmp, list)
			result = append(result, tmp)
		}

		for i := start; i < len(s); i++ {
			if isPartition(s[start : i+1]) {
				list = append(list, s[start:i+1])
			} else {
				continue
			}
			dfs(i + 1)
			list = list[:len(list)-1]
		}
	}
	dfs(0)

	return result
}
