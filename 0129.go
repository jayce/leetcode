/*
 * Creation in 2020-11-10 by Chuan Liu (Jayce)
 */

package leetcode

// sumNumbers First idea
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0) // time: O(N)
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}
