/*
 * Creation in 2020-11-16 by Chuan Liu (Jayce)
 */

package leetcode

// searchBST First idea
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}

// Second idea
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			break
		}
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}
