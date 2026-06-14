/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {

	// [1,2,3] -> [2,3] -> 2 then 3. 3 becomes latest of that cycle (right node)
	// trick is by adding left first so right can become last node added in the cycle

	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	qLen := 0
	res := []int{}

	for len(q) != 0 {

		qLen = len(q)
		var latest *TreeNode
		for i:=0; i<qLen; i++ {
			latest = q[0] // update each cycle
			q = q[1:]

			if latest.Left != nil {
				q = append(q, latest.Left)
			}

			if latest.Right != nil {
				q = append(q, latest.Right)
			}
		}
		res = append(res, latest.Val)
	}
	return res

}
