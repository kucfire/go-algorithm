/*
	leetcode tag:524 title:通过删除字母匹配到字典里最长单词
*/

package point

// method of double points
func findLongestWord(s string, dictionary []string) string {
	result := ""
	for _, word := range dictionary {
		if subString(s, word) {
			if (len(result) < len(word)) || (len(result) == len(word) && word < result) {
				result = word
			}
		}
	}
	return result
}

func subString(s, target string) bool {
	if len(target) <= 0 {
		return true
	}
	index1, index2 := 0, 0
	for index1 < len(s) {
		if s[index1] == target[index2] {
			index2++
		}
		index1++
		if index2 == len(target) {
			return true
		}
	}

	return false
}
