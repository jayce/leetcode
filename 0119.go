/*
 * Creation in 2020-11-20 by Chuan Liu (Jayce)
 */

package leetcode

// rightSideView First idea
func rightSideView(root *TreeNode) []int {
	data := []int{}
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) { // time: O(N)
		if root == nil {
			return
		}
		if len(data) == depth {
			data = append(data, root.Val)
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	dfs(root, 0)
	return data
}
