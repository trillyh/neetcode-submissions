/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		var latest *TreeNode
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			latest = queue[0]
			queue = queue[1:]

			if latest.Left != nil {
				queue = append(queue, latest.Left)
			}
			if latest.Right != nil {
				queue = append(queue, latest.Right)
			}	
		}
		res = append(res, latest.Val)
	}
	return res
    
}
