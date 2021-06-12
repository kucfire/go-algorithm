/*
	leetcode tag:452 title:用最少数量的箭引爆气球
*/

package greedy

import (
	"fmt"
	"sort"
)

// method of greedy
func FindMinArrowShots(points [][]int) int {
	result := 0
	if len(points) == 0 {
		return result
	}

	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	fmt.Println(points)
	right := points[0][1]
	result++
	for _, point := range points[1:] {
		if right < point[0] {
			result++
			right = point[1]
		}
	}

	return result
}
