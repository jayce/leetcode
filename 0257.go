/*
 * Creation in 2020-11-24 by Chuan Liu (Jayce)
 */

package leetcode

import "fmt"

// binaryTreePaths First idea
func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) { // time: O(N)
		if root == nil {
			return
		}
		path = fmt.Sprintf("%s%d", path, root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, path)
			return
		}
		dfs(root.Left, path+"->")
		dfs(root.Right, path+"->")
	}
	dfs(root, "")
	return paths
}
