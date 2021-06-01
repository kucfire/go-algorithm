/*
	leetcode tag:279 title:完全平方数
*/

package bfs

// failed
func numSquaresMyself(n int) int {
	result := 0
	for i := 0; i <= n && n > 0; i++ {
		if i*i >= n {
			result++
			n = n - (i-1)*(i-1)
			i = 0
		}
	}

	return result
}

// method of BFS
func numSquares(n int) int {
	queue := make([]int, 0)
	squares := generateSquares(n)
	marked := make([]bool, n+1)

	queue = append(queue, n)
	marked[n] = true

	level := 0

	for len(queue) > 0 {
		size := len(queue)
		level++
		for ; size > 0; size-- {
			cur := queue[0]
			queue = queue[1:]
			for _, s := range squares {
				next := cur - s
				if next < 0 {
					break
				}
				if next == 0 {
					return level
				}
				if marked[next] {
					continue
				}
				marked[next] = true
				queue = append(queue, next)
			}
		}
	}

	return n
}

func generateSquares(n int) []int {
	result := make([]int, 0)
	for i := 1; i <= n; i++ {
		result = append(result, i*i)
	}
	return result
}
