/*
	leetcode tag:017 title:电话号码的字母组合
*/

package backtracking

// method of backtracking
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := make([]string, 0)

	dictionary := map[rune][]string{
		'1': {""},
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	getResult := func(a, b []string) []string {
		newResult := make([]string, 0)

		for _, strA := range a {
			for _, strB := range b {
				newResult = append(newResult, strA+strB)
			}
		}
		return newResult
	}

	for _, str := range digits {
		dic := dictionary[str]
		if len(result) == 0 {
			result = append(result, dic...)
			continue
		}
		result = getResult(result, dic)
	}

	return result
}
