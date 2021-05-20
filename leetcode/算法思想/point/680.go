/*
	leetcode tag:680 title:验证回文字符串 Ⅱ
*/

package point

// method of double points
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			flag1, flag2 := true, true
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

// false
func validPalindromeMyself(s string) bool {
	delSign := false
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			if delSign {
				return false
			}
			if s[left+1] == s[right] {
				left += 2
				right--
				delSign = true
				continue
			}
			if s[left] == s[right-1] {
				left++
				right -= 2
				delSign = true
				continue
			}
			return false
		}
	}
	return true
}
