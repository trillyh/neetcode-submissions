/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var isValid func(*TreeNode, int64, int64) bool

	isValid = func(node *TreeNode, min int64, max int64) bool {
		if node == nil {
			return true
		}

		val := int64(node.Val)
		left := isValid(node.Left, min, val)
		right := isValid(node.Right, val, max)

		return left && right && min < val && val < max
	} 
	return isValid(root, math.MinInt64, math.MaxInt64)
}
