/*
	leetcode tag:257 title:二叉树的所有路径
*/

package backtracking

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// method of backtracking
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)

	dfs := func(point *TreeNode, route []string) {}
	dfs = func(point *TreeNode, route []string) {
		if point == nil {
			return
		}
		value := strconv.Itoa(point.Val)
		route = append(route, value)
		if point.Left == nil && point.Right == nil {
			if len(route) > 0 {
				tmp := strings.Join(route, "->")
				result = append(result, tmp)
			}
			return
		}
		dfs(point.Left, route)
		dfs(point.Right, route)
	}

	dfs(root, []string{})
	return result
}
