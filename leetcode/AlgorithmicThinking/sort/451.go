/*
	leetcode tag:451 title:根据字符出现频率排序
*/

package sort

import "math"

// 桶排序
func frequencySort(s string) string {
	hashmap := make(map[rune]int)
	maxn := math.MinInt64
	for _, str := range s {
		if _, ok := hashmap[str]; ok {
			hashmap[str]++
		} else {
			hashmap[str] = 1
		}
		if hashmap[str] > maxn {
			maxn = hashmap[str]
		}
	}

	hashtop := make([][]rune, maxn+1)
	for key, value := range hashmap {
		hashtop[value] = append(hashtop[value], key)
	}

	result := make([]rune, 0)
	for i := maxn; i > 0; i-- {
		for _, str := range hashtop[i] {
			for j := i; j > 0; j-- {
				result = append(result, str)
			}
		}
	}
	return string(result)
}
