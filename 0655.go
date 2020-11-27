/*
 * Creation in 2020-11-27 by Chuan Liu (Jayce)
 */

package leetcode

import "strconv"

// printTree First idea
func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int { // time: O(N)
		if root == nil {
			return 0
		}
		return max(dfs(root.Left), dfs(root.Right)) + 1
	}
	depth := dfs(root)

	nums := 0
	for i := 0; i < depth; i++ { // time: O(H)
		nums += 1 * i
	}

	data := make([][]string, depth)
	for i := 0; i < depth; i++ { // time: O(H)
		data[i] = make([]string, nums)
	}

	var fill func(root *TreeNode, depth, s, e int)
	fill = func(root *TreeNode, depth, s, e int) { // time: O(N)
		if root == nil || s > e {
			return
		}
		m := s + (e-s)/2
		data[depth][m] = strconv.Itoa(root.Val)
		fill(root.Left, depth+1, s, m-1)
		fill(root.Left, depth+1, m+1, e)
	}
	fill(root, 0, 0, nums)
	return data
}
