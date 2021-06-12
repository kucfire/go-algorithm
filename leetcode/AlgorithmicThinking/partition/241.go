/*
	leetcode tag:241 title:为运算表达式设计优先级
*/

package partition

import "strconv"

// method of partition
func diffWaysToCompute(expression string) []int {
	if isNum(expression) {
		tmp, _ := strconv.Atoi(expression)
		return []int{tmp}
	}

	var res []int
	for i, c := range expression {
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
			// 如果是运算符，则计算左右两边的等式
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int

					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else if tmpC == "*" {
						addNum = leftNum * rightNum
					}

					res = append(res, addNum)
				}
			}
		}
	}
	return res
}

func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}
