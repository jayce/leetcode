/*
 * Creation in 2020-11-23 by Chuan Liu (Jayce)
 */

package leetcode

// insertIntoBST First idea bottom insertion
func insertIntoBST(root *TreeNode, val int) *TreeNode { // time: O(logN)
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val) // time: O(N/2)
	} else {
		root.Right = insertIntoBST(root.Right, val) // time: O(N/2)
	}
	return root
}

// Second idea
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	p := root
	for p != nil { // time: O(logN)
		if p.Val < val {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				break
			}
			p = p.Right
		} else {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				break
			}
			p = p.Left
		}
	}
	return root
}
