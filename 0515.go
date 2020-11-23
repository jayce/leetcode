/*
 * Creation in 2020-11-23 by Chuan Liu (Jayce)
 */

package leetcode

// largestValues First idea from traversal of level
func largestValues(root *TreeNode) []int {
	largests := make([]int, 0, 1024)
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) { // time: O(N)
		if root == nil {
			return
		}
		if len(largests) == depth {
			largests = append(largests, root.Val)
		}
		if largests[depth] < root.Val {
			largests[depth] = root.Val
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
		return
	}
	dfs(root, 0)
	return largests
}
