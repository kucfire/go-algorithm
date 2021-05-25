/*
	leetcode tag:605 title:种花问题
*/

package greedy

// method of greedy
func canPlaceFlowers(flowerbed []int, n int) bool {
	total := 0
	pre := -1
	for i := range flowerbed {
		if flowerbed[i] == 1 {
			if pre < 0 {
				total += i / 2
			} else {
				total += (i - pre - 2) / 2 // 这里-2，是需要拆分成两个-1来看，首先i-pre得出来的结果是i与pre的差值，需要-1来得出i与pre的中间间隔；另一个-1是需要减去中间间隔的无效值，该无效值会影响/2产生的结果
			}
			pre = i
		}
	}

	if pre < 0 {
		total += (len(flowerbed) + 1) / 2
	} else {
		total += (len(flowerbed) - pre - 1) / 2
	}

	return total >= n
}
