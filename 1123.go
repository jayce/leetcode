/*
 * Creation in 2021-05-19 by Chuan Liu (Jayce)
 */

package leetcode

// First idea, inspired by the problem 236.
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	node, _ := dfs(root) // time: O(N)
	return node
}

func dfs(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return root, 0
	}

	Lnode, Ldepth := dfs(root.Left)
	Rnode, Rdepth := dfs(root.Right)
	if Ldepth == Rdepth {
		return root, Rdepth + 1
	}
	if Ldepth > Rdepth {
		return Lnode, Ldepth + 1
	}
	return Rnode, Rdepth + 1
}
