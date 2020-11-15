/*
 * Creation in 2020-11-13 by Chuan Liu (Jayce)
 */

package leetcode

// First idea
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > q.Val { // swap p,q make p < q
		return dfs(root, q, p) // time: O(N)
	}
	return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > root.Val { // p > root => root.Right
		return lowestCommonAncestor(root.Right, p, q)
	}
	if q.Val < root.Val { // q < root => root.Left
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}
