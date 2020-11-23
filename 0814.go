/*
 * Creation in 2020-11-23 by Chuan Liu (Jayce)
 */

package leetcode

// pruneTree First idea
func pruneTree(root *TreeNode) *TreeNode {
	if dfs(root) == 0 {
		return nil
	}
	return root
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	if left == 0 {
		root.Left = nil
	}
	right := dfs(root.Right)
	if right == 0 {
		root.Right = nil
	}
	return left + right + root.Val
}
