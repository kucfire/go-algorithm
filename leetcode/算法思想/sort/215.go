/*
	leetcode tag:215 title:数组中的第K个最大元素
*/

package sort

import (
	"math/rand"
	"sort"
	"time"
)

// method of sort
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 堆（看不懂，我也在努力理解中）
// func findKthLargestHeap(nums []int, k int) int {
// 	lens := len(nums) - 1
// 	for i := lens >> 1; i >= 0; i-- {
// 		down(nums, i, lens)
// 	}
// 	for j := lens; j >= 1; j-- {
// 		nums[0], nums[j] = nums[j], nums[0]
// 		lens--
// 		down(nums, 0, lens)
// 	}
// 	return nums[k-1]
// }

// func down(nums []int, i, lens int) {
// 	min := i
// 	if i<<1+1 <= lens && nums[i<<1+1] < nums[min] { //相等最后一个元素
// 		min = i<<1 + 1
// 	}
// 	if i<<1+2 <= lens && nums[i<<1+2] < nums[min] {
// 		min = i<<1 + 2
// 	}
// 	if i != min {
// 		nums[i], nums[min] = nums[min], nums[i]
// 		down(nums, min, lens)
// 	}
// }

// 堆
func findKthLargestHeap(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, i, heapSize)
	}

	return nums[0]
}
func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}
func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

// func findKthLargestHeap(nums []int, k int) int {
// 	heapSize := len(nums)
// 	buildMaxHeap(nums, heapSize)
// 	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
// 		nums[0], nums[i] = nums[i], nums[0]
// 		heapSize--
// 		maxHeapify(nums, 0, heapSize)
// 	}
// 	return nums[0]
// }

// func buildMaxHeap(a []int, heapSize int) {
// 	for i := heapSize / 2; i >= 0; i-- {
// 		maxHeapify(a, i, heapSize)
// 	}
// }

// func maxHeapify(a []int, i, heapSize int) {
// 	l, r, largest := i*2+1, i*2+2, i
// 	if l < heapSize && a[l] > a[largest] {
// 		largest = l
// 	}
// 	if r < heapSize && a[r] > a[largest] {
// 		largest = r
// 	}
// 	if largest != i {
// 		a[i], a[largest] = a[largest], a[i]
// 		maxHeapify(a, largest, heapSize)
// 	}
// }

// method of fast sort
func findKthLargestFastSort(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, index int) int {
	q := randomPartition(nums, left, right)
	if q == index {
		return nums[q]
	} else if q < index {
		return quickSelect(nums, q+1, right, index)
	}
	return quickSelect(nums, left, q-1, index)
}

func randomPartition(nums []int, left, right int) int {
	i := rand.Int()%(right-left+1) + left
	nums[i], nums[right] = nums[right], nums[i]
	return partition(nums, left, right)
}
func partition(nums []int, left, right int) int {
	x := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}
