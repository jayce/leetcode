/*
 * Creation in 2020-11-16 by Chuan Liu (Jayce)
 */

package leetcode

// isBalanced First idea
func isBalanced(root *TreeNode) bool {
	b := true
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int { // time: O(N)
		if root == nil {
			return 0
		}
		L := dfs(root.Left)
		R := dfs(root.Right)
		if abs(L-R) > 1 {
			b = false
		}
		return max(L, R) + 1
	}
	dfs(root)
	return b
}

// Second idea optimized for the First
func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		L := dfs(root.Left)
		if L == -1 {
			return L
		}
		R := dfs(root.Right)
		if R == -1 {
			return R
		}
		if abs(L-R) > 1 {
			return -1
		}
		return max(L, R) + 1
	}
	return dfs(root) != -1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
