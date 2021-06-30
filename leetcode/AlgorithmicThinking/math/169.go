/*
	leetcode tag:169 title:多数元素
*/

package math

// method of math
func majorityElement(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	for k, v := range count {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

// method of 摩尔投票
func majorityElement2(nums []int) int {
	count, target := 0, 0
	for _, num := range nums {
		if count == 0 {
			target = num
			count++
			continue
		}
		if target == num {
			count++
		} else {
			count--
		}
	}
	return target
}
