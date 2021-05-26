/*
	leetcode tag:665 title:非递减数列
*/

package greedy

// method of greedy
func CheckPossibility(nums []int) bool {
	signed := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if !signed {
				signed = true
			} else {
				return false
			}
			if i == 1 || (i > 1 && nums[i-2] < nums[i]) {
				nums[i-1] = nums[i]
			}
			if i > 1 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			}
		}
	}
	return true
}
