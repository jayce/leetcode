/*
 * Creation in 2020-11-23 by Chuan Liu (Jayce)
 */

package leetcode

// diameterOfBinaryTree First idea from No.104
func diameterOfBinaryTree(root *TreeNode) int {
	maximumPath := 0
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int { // time: O(N)
		if root == nil {
			return 0
		}
		L := depth(root.Left)
		R := depth(root.Right)
		maximumPath = max(maximumPath, L+R)
		return max(L, R) + 1
	}
	depth(root)
	return maximumPath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
