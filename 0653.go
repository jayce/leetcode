/*
 * Creation in 2020-11-17 by Chuan Liu (Jayce)
 */

package leetcode

// findTarget First idea
func findTarget(root *TreeNode, k int) bool {
	hash := make(map[int]bool) // space: O(N)
	var dfs func(root *TreeNode, k int) bool
	dfs = func(root *TreeNode, k int) bool { // time: O(N)
		if root == nil {
			return false
		}
		if hash[k-root.Val] {
			return true
		}
		hash[root.Val] = true
		return dfs(root.Left, k) || dfs(root.Right, k)
	}
	return dfs(root, k)
}

// Second idea
func findTarget(root *TreeNode, k int) bool {
	hash := make(map[int]bool)          // space: O(N)
	stack := make([]*TreeNode, 0)       // space: O(N)
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if hash[k-root.Val] {
			return true
		}

		hash[root.Val] = true
		root = root.Right
	}
	return false
}
