/*
	leetcode tag:075 title:颜色分类
*/

package sort

// 桶排序
func sortColors(nums []int) {
	top := make([][]int, 3)
	for _, num := range nums {
		top[num] = append(top[num], num)
	}

	i := 0
	for _, top2 := range top {
		for _, element := range top2 {
			nums[i] = element
			i++
		}
	}
}
