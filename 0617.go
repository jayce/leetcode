/*
 * Creation in 2020-11-15 by Chuan Liu (Jayce)
 */

package leetcode

// mergeTrees First idea
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode { // time: O(N)
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
