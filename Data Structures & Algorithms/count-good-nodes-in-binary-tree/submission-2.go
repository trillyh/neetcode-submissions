/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	count := 0
	max := root.Val
	var dfs func(*TreeNode, int) 	
	dfs = func(node *TreeNode, max int) {
		if node == nil {
			return
		}					

		if node.Val >= max {
			count++
			max = node.Val
		}

		dfs(node.Left, max)
		dfs(node.Right, max)
	}
	dfs(root, max)
	return count
}

func intMax(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}