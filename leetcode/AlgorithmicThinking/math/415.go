/*
	leetcode tag:415 title:字符串想加
*/

package math

import "strconv"

// method of math
func addStrings(num1 string, num2 string) string {
	add := 0
	result := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		tmp := 0
		if i >= 0 {
			tmp += int(num1[i] - '0')
		}
		if j >= 0 {
			tmp += int(num2[j] - '0')
		}
		if add != 0 {
			tmp += add
		}
		result = strconv.Itoa(tmp%10) + result
		add = tmp / 10
	}

	return result
}
