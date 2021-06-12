/*
	leetcode tag:347 title:前K个高频元素
*/

package sort

import "math"

// 桶排序
func topKFrequent(nums []int, k int) []int {
	hashkey := make(map[int]int)
	maxn := math.MinInt64
	for _, num := range nums {
		if _, ok := hashkey[num]; ok {
			hashkey[num]++
		} else {
			hashkey[num] = 1
		}
		if hashkey[num] > maxn {
			maxn = hashkey[num]
		}
	}

	hashtop := make([][]int, maxn+1)
	for key, value := range hashkey {
		hashtop[value] = append(hashtop[value], key)
	}

	res := make([]int, 0)
	for i := maxn; i >= 0; i-- {
		res = append(res, hashtop[i]...)
		k -= len(hashtop[i])
		if k == 0 {
			break
		}
	}
	return res
}
