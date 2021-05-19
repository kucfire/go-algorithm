/*
	leetcode tag:345 title:反转字符串中的元音字母
*/

package leetcode

// method of double points
func reverseVowels(s string) string {
	b := []byte(s)
	// 先定一个hashmap
	hashmap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	left, right := 0, len(s)-1
	for left < right {
		vl := hashmap[b[left]]
		vr := hashmap[b[right]]
		if vl && vr {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
		if !vl {
			left++
		}
		if !vr {
			right--
		}
	}

	return string(b)
}
