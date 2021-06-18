package main

import (
	"fmt"
	dynamicprogramming "go-algorithm/leetcode/AlgorithmicThinking/DynamicProgramming"
)

// test case
func main() {
	// 435
	// fmt.Println(greedy.EraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	// fmt.Println(greedy.EraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}))
	// fmt.Println(greedy.EraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}))
	// 452
	// fmt.Println(greedy.FindMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	// fmt.Println(greedy.FindMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	// 406
	// fmt.Println(greedy.ReconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	// 665
	// fmt.Println(greedy.CheckPossibility([]int{3, 4, 2, 3}))
	// fmt.Println(greedy.CheckPossibility([]int{5, 7, 1, 8}))
	// 744
	// fmt.Println(binarysearch.NextGreatestLetter([]byte{'c', 'f', 'j'}, byte('c')))
	// 540
	// fmt.Println(binarysearch.SingleNonDuplicate([]int{1}))

	// binarysearch.SingleNonDuplicate
	fmt.Println(dynamicprogramming.LongestCommonSubsequence("abcde", "ace"))
}
