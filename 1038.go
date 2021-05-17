/*
 * Creation in 2021-05-17 by Chuan Liu (Jayce)
 */

package leetcode

// First idea, from inoder
func bstToGst(root *TreeNode) *TreeNode {
	last := (*TreeNode)(nil)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) { // time: O(N)
		if root == nil {
			return
		}

		inorder(root.Right)
		if last != nil {
			root.Val += last.Val
		}
		last = root
		inorder(root.Left)
	}

	inorder(root)
	return root
}

// Second idea
func bstToGst(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0, 128) // space: O(H)
	last, p := (*TreeNode)(nil), root
	for len(stack) > 0 || p != nil { // time: O(N)
		for p != nil {
			stack = append(stack, p)
			p = p.Right
		}

		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last != nil {
			p.Val += last.Val
		}
		last = p
		p = p.Left
	}

	return root
}
