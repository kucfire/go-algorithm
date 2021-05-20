/*
	leetcode tag:167 title:两数之和 II - 输入有序数组
*/

package point

// method of loop
func twoSum167(numbers []int, target int) []int {
	for first := 0; first < len(numbers)-1; first++ {
		for second := first + 1; second < len(numbers); second++ {
			if numbers[first]+numbers[second] == target {
				return []int{first + 1, second + 1}
			}
		}
	}
	return nil
}

// method of hashmap(best)
func twoSum167_2(numbers []int, target int) []int {
	hashmap := make(map[int]int)
	for i, num := range numbers {
		if v, ok := hashmap[num]; ok {
			return []int{v + 1, i + 1}
		}
		hashmap[target-num] = i
	}
	return nil
}

// method of double points
func twoSum167_3(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return nil
}
