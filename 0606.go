/*
 * Creation in 2020-11-25 by Chuan Liu (Jayce)
 */

package leetcode

import "strconv"

// tree2str First idea
func tree2str(t *TreeNode) string { // time: O(N)
	if t == nil {
		return ""
	}
	s := strconv.Itoa(t.Val)
	L := tree2str(t.Left)
	R := tree2str(t.Right)
	if L != "" || R != "" {
		s = s + "(" + L + ")"
	}
	if R != "" {
		s = s + "(" + R + ")"
	}
	return s
}
