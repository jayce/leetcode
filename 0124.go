/*
 * Creation in 2020-11-10 by Chuan Liu (Jayce)
 */

package leetcode

// maxPathSum First idea
func maxPathSum(root *TreeNode) int {
	maximumPathSum := ^int(^uint(0) >> 1) // init: minimum
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lPathSum := dfs(root.Left)
		rPathSum := dfs(root.Right)
		maximumPathSum = max(lPathSum+rPathSum+root.Val, maximumPathSum)
		return max(max(lPathSum, rPathSum)+root.Val, 0)
	}
	dfs(root)
	return maximumPathSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
