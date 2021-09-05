package arrayandmatrix

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func duplicate(numbers []int) int {
	// write code here
	hash := make(map[int]bool) // 定义一个map用来存储已出现过的数字

	for _, num := range numbers {
		if hash[num] == true { // 当num已经在map里面时，证明已重复，则直接return该num
			return num
		} else { // num不存在map中，则将num存进map（将value的bool设置为true
			hash[num] = true
		}
	}

	return -1
}
