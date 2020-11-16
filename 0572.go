/*
 * Creation in 2020-11-16 by Chuan Liu (Jayce)
 */

package leetcode

import (
	"strconv"
	"strings"
)

// isSubtree First idea time: O(s*t)
func isSubtree(s *TreeNode, t *TreeNode) bool {
	return s != nil && (dfs(s, t) || // time: O(t)
		isSubtree(s.Left, t) || // time: O(s)
		isSubtree(s.Right, t))
}

func dfs(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	return s.Val == t.Val &&
		dfs(s.Left, t.Left) &&
		dfs(s.Right, t.Right)
}

// transfer to the strings
func isSubtree(s *TreeNode, t *TreeNode) bool {
	var dfs func(s *TreeNode) string
	dfs = func(s *TreeNode) string {
		if s == nil {
			return "$"
		}
		return "[" + strconv.Itoa(s.Val) + "]" + dfs(s.Left) + dfs(s.Right)
	}
	return strings.Contains(dfs(s), dfs(t))
}
