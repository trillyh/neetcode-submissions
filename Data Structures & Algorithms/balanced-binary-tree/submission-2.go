/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Value struct { 
	balance bool
	depth int
}

func dfs(root *TreeNode) Value {
	if root == nil {
		return Value{balance: true, depth: 0}
	}	

	left := dfs(root.Left)
	right := dfs(root.Right)

	isBalance := int(math.Abs(float64(left.depth)-float64(right.depth))) <= 1
	isBalance = isBalance && left.balance && right.balance

	maxDepth := max(left.depth, right.depth) + 1
	return Value{balance: isBalance, depth: maxDepth}
}

func isBalanced(root *TreeNode) bool {
	return dfs(root).balance 
}
