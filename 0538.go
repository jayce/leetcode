/*
 * Creation in 2021-05-17 by Chuan Liu (Jayce)
 */

package leetcode

func convertBST(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0, 512)
	last, p := (*TreeNode)(nil), root
	for len(stack) > 0 || p != nil {
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
