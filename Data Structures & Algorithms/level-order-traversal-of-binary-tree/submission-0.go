/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	
	if root == nil {
		return res
	}
	
	q := []*TreeNode{root}

	for len(q) > 0 {
		qLen := len(q)
		currLevel := []int{}

		for i := 0; i < qLen; i++ {
			node := q[0]
			currLevel = append(currLevel, node.Val)

			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)	
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, currLevel)
	}
	return res 
}
