/*
	leetcode tag:763 title:划分字母区间
*/

package greedy

// method of greedys
func partitionLabels(s string) []int {
	list := make([]int, 26)
	for i, str := range s {
		list[str-'a'] = i
	}

	result := make([]int, 0)
	start, end := 0, 0
	for i, str := range s {
		if list[str-'a'] > end { // 找出区间内的最大下标
			end = list[str-'a']
		}

		if end == i { // 当i遍历到与end相同时，代表该区间内没有其他更大的下标
			result = append(result, end-start+1)
			start = end + 1
		}
	}

	return result
}
