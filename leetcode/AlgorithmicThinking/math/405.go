/*
	leetcode tag:405 title:数字转换为十六进制数
*/

package math

import (
	"strings"
)

// method of math
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	num = num & 0xffffffff
	result := make([]string, 0)
	list := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f",
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
		result = append(result, list[num%16])
		num = num / 16
	}

	reverse(result)
	return strings.Join(result, "")
}
