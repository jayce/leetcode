/*
 * Creation in 2020-11-15 by Chuan Liu (Jayce)
 */

package leetcode

// findTilt First idea
func findTilt(root *TreeNode) int {
	tilt := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int { // time: O(N)
		if root == nil {
			return 0
		}
		L, R := dfs(root.Left), dfs(root.Right)
		tilt += abs(L - R)
		return L + R + root.Val
	}
	dfs(root)
	return tilt
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
