/*
 * Creation in 2020-11-15 by Chuan Liu (Jayce)
 */

// isUnivalTree First idea
func isUnivalTree(root *TreeNode) bool { // time: O(N)
	if root == nil {
		return root == nil
	}

	if root.Left != nil && root.Val != root.Left.Val {
		return false
	}

	if !isUnivalTree(root.Left) {
		return false
	}

	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}
	return isUnivalTree(root.Right)
}