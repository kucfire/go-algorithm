/*
	leetcode tag:435 title:无重叠区间
*/

package greedy

import (
	"fmt"
	"sort"
)

// method of greedy
func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	fmt.Println(intervals)
	result, right := 1, intervals[0][1]

	for _, interval := range intervals[1:] {
		if interval[0] >= right {
			result++
			right = interval[1]
		}
	}

	return n - result
}
