/*
	leetcode tag:406 title:根据身高重建队列
*/
package greedy

import (
	"fmt"
	"sort"
)

// method of greedy
func ReconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	fmt.Println(people)

	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]] = p
	}

	return people
}
