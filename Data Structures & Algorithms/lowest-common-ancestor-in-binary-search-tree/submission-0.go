/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // if p < node < q -> root is LCA and returns.
 // if p == node || q == node -> root is LCA and returns.
 // if node < p and q -> go right
 // if node > p and q -> go left
 // 

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    if p.Val > q.Val {
        p, q = q, p
    }
    var lca *TreeNode
    var dfs func(*TreeNode, *TreeNode, *TreeNode)
    dfs = func(node *TreeNode, p *TreeNode, q *TreeNode){
        if node == nil {
            return
        }
        fmt.Println("node is %d %t", node.Val, (p.Val < node.Val && node.Val < q.Val) )
        if (p.Val < node.Val && node.Val < q.Val) ||
        (node.Val == p.Val || node.Val == q.Val) {
            lca = node
            fmt.Println("node is", node.Val)
            return
        }
        if node.Val < p.Val && node.Val < q.Val {
            dfs(node.Right, p, q)
        } else {
            dfs(node.Left, p, q)
        }
    }
    dfs(root, p, q)
    return lca
}
