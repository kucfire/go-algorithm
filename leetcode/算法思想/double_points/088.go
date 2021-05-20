/*
	leetcode tag:088 title:合并两个有序数组
*/

package point

import "sort"

// method of double points
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	indexMerge := m + n - 1
	for index2 >= 0 {
		if index1 < 0 {
			nums1[indexMerge] = nums2[index2]
			indexMerge--
			index2--
			continue
		}
		if nums1[index1] > nums2[index2] {
			nums1[indexMerge] = nums1[index1]
			indexMerge--
			index1--
			continue
		}
		if nums1[index1] < nums2[index2] {
			nums1[indexMerge] = nums2[index2]
			indexMerge--
			index2--
			continue
		}
		if nums1[index1] == nums2[index2] {
			nums1[indexMerge] = nums1[index1]
			nums1[indexMerge-1] = nums2[index2]
			index1--
			index2--
			indexMerge -= 2
			continue
		}
	}
}

// most simple
func mergeSort(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
