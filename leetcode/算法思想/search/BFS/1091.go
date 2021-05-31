/*
	leetcode tag:1091 title:二进制矩阵中的最短路径
*/

package bfs

// method of bfs广度优先搜索
func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) <= 0 || grid[0][0] == 1 {
		return -1
	}
	list := [][]int{{0, 0, 1}} // 左边的位置 右边的位置 长度
	// result := []int{}
	for len(list) > 0 {
		element := list[0]
		list = list[1:]
		// 校验是不是到达右下的单元格，如果是的话，返回element的长度，第一个遍历到证明就是最短路径
		if element[0] == len(grid)-1 && element[1] == len(grid[0])-1 {
			return element[2]
		}
		// 先校验可能的最短路径的点，即横纵坐标都+1的点
		if element[0]+1 <= len(grid)-1 && element[1]+1 <= len(grid[0])-1 {
			if grid[element[0]+1][element[1]+1] == 0 {
				list = append(list, []int{element[0] + 1, element[1] + 1, element[2] + 1})
				grid[element[0]+1][element[1]+1] = 1
			}
		}
		// 校验可能的最短路径点的左和上两个临近当前点的点
		if element[0]+1 <= len(grid)-1 {
			if grid[element[0]+1][element[1]] == 0 {
				list = append(list, []int{element[0] + 1, element[1], element[2] + 1})
				grid[element[0]+1][element[1]] = 1
			}
		}
		if element[1]+1 <= len(grid[0])-1 {
			if grid[element[0]][element[1]+1] == 0 {
				list = append(list, []int{element[0], element[1] + 1, element[2] + 1})
				grid[element[0]][element[1]+1] = 1
			}
		}
		// 校验最后的可能点
		if element[0]+1 <= len(grid)-1 && element[1]-1 >= 0 {
			if grid[element[0]+1][element[1]-1] == 0 {
				list = append(list, []int{element[0] + 1, element[1] - 1, element[2] + 1})
				grid[element[0]+1][element[1]-1] = 1
			}
		}
		if element[1]+1 <= len(grid[0])-1 && element[0]-1 >= 0 {
			if grid[element[0]-1][element[1]+1] == 0 {
				list = append(list, []int{element[0] - 1, element[1] + 1, element[2] + 1})
				grid[element[0]-1][element[1]+1] = 1
			}
		}
	}

	return -1
}

/*
[0,1,1,0,0,0],
[0,1,0,1,1,0],
[0,1,1,0,1,0],
[0,0,0,1,1,0],
[1,1,1,1,1,0],
[1,1,1,1,1,0]]
*/
