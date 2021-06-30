/*
	leetcode tag:504 title:七进制数
*/

package math

import (
	"strconv"
	"strings"
)

// method of math
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	result := make([]string, 0)
	mark := false
	if num < 0 {
		mark = true
		num = -num
	}

	reverse := func(list []string) {
		left, right := 0, len(list)-1
		for left <= right {
			list[left], list[right] = list[right], list[left]
			left++
			right--
		}
	}

	for num != 0 {
		result = append(result, strconv.Itoa(num%7))
		num = num / 7
	}

	if mark {
		result = append(result, "-")
	}

	reverse(result)

	return strings.Join(result, "")
}
