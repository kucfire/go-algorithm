/*
	leetcode tag:168 title:Excel表列名称
*/

package math

import "strings"

// method of math
func convertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}
	list := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	result := make([]string, 0)
	reverse := func(list []string) {
		left, right := 0, len(list)-1
		for left <= right {
			list[left], list[right] = list[right], list[left]
			left++
			right--
		}
	}

	for columnNumber != 0 {
		columnNumber--
		result = append(result, list[columnNumber%26])
		columnNumber = columnNumber / 26
	}
	reverse(result)

	return strings.Join(result, "")
}
