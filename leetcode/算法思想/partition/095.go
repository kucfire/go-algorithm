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
	return
}

func getResult(start, end int) []*TreeNode {

}
