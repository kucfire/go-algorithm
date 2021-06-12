/*
	leetcode tag:744 title:寻找比目标字母大的最小字母
*/

package binarysearch

// method of binary search
func NextGreatestLetter(letters []byte, target byte) byte {
	result := letters[len(letters)-1]

	if target >= result {
		return letters[0]
	}

	l, r := 0, len(letters)-2

	for l <= r {
		mid := l + (r-l)/2
		if letters[mid] > target {
			if letters[mid] < result {
				result = letters[mid]
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result
}
