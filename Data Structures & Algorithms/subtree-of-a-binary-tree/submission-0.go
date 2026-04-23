/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	left := isSubtree(root.Left, subRoot)
	right := isSubtree(root.Right, subRoot)

	isCurrSubtree := false
	if root.Val == subRoot.Val {
		isCurrSubtree = sameTree(root, subRoot)
	}

	return left || right || isCurrSubtree

    
}

func sameTree(root1 *TreeNode, root2 *TreeNode) bool {
	// base case: if both are nil, return true <- pass this up
	// if they are equal in value, and both are not nil, return true
	// case they fail is when either are nil but not both, AND they not equal in value
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 != nil && root1.Val == root2.Val {
		return sameTree(root1.Left, root2.Left) && sameTree(root1.Right, root2.Right)
	}
	return false
}
