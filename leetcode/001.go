/*
	leetcode tag:001 title:两数之和
*/
package leetcode

// method of hashmap(best)
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, num := range nums {
		if v, ok := hashmap[num]; ok {
			return []int{v, i}
		}
		hashmap[target-num] = i
	}
	return nil
}

// method of double points
func twoSum2(nums []int, target int) []int {
	for first := 0; first < len(nums)-1; first++ {
		for second := first + 1; second < len(nums); second++ {
			if nums[first]+nums[second] == target {
				return []int{first, second}
			}
		}
	}
	return nil
}
