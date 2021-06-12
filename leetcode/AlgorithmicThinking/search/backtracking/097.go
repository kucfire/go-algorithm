/*
	leetcode tag:097 title:复原IP地址
*/

package backtracking

import (
	"strconv"
	"strings"
)

// method of backtracking
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 32 {
		return []string{}
	}

	result := make([]string, 0)
	track := make([]string, 0)
	// result = append(result, backtrack(s, track, result, 1))
	backtrack(s, track, &result, 1)
	return result
}

func backtrack(s string, track []string, result *[]string, key int) {
	// key为段数，一共四段，当key为5时证明已经选完四段了，同时字符串选完时结束
	if key == 5 {
		if len(s) == 0 {
			str := strings.Join(track, ".")
			*result = append(*result, str)
		}
	}

	// 选择列表
	// 每一段最大选择三位
	for i := 1; i <= 3; i++ {
		if i <= len(s) {
			// 选1-3位数
			v, err := strconv.Atoi(s[:i])
			if err == nil && v <= 255 {
				// 做选择
				//
				track = append(track, s[:i])
				str := s[i:]
				// 下一段选择
				backtrack(str, track, result, key+1)
				// 撤销选择
				track = track[:len(track)-1]
			}
			if v == 0 {
				break
			}
		}
	}
}
