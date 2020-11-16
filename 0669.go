/*
 * Creation in 2020-11-16 by Chuan Liu (Jayce)
 */

package leetcode

// trimBST First idea
func trimBST(root *TreeNode, low int, high int) *TreeNode { // time: O(N)
	if root == nil {
		return root
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	if root.Val < low {
		return root.Right
	}
	if root.Val > high {
		return root.Left
	}
	return root
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// trim root first
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		}
		if root.Val > high {
			root = root.Left
		}
	}

	// trim left subtree
	dummy := root
	for dummy != nil {
		for dummy.Left != nil && dummy.Left.Val < low {
			dummy.Left = dummy.Left.Right
		}
		dummy = dummy.Left
	}
	dummy = root
	for dummy != nil {
		for dummy.Right != nil && dummy.Right.Val > high {
			dummy.Right = dummy.Right.Left
		}
		dummy = dummy.Right
	}
	return root
}
