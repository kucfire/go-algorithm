/*
	leetcode tag:095 title:不同的二叉搜索树 II
*/

package partition

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// method of partition
func generateTrees(n int) []*TreeNode {
	// 二叉搜索树，需要满足从左到右，从小到大
	if n == 0 {
		return nil
	}
	return getResult(1, n)
}

func getResult(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	allTrees := []*TreeNode{}

	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTree := getResult(start, i-1)
		// 获得所有可行的右子树集合
		rightTree := getResult(i+1, end)
		// 从左子树集合中选出一颗左子树，从右子树结合中选出一颗右子树，拼接到节点上
		for _, left := range leftTree {
			for _, right := range rightTree {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}

	return allTrees
}
