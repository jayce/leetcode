/*
 * Creation in 2020-11-25 by Chuan Liu (Jayce)
 */

package leetcode

// invertTree First idea
func invertTree(root *TreeNode) *TreeNode { // time: O(N)
	if root == nil {
		return root
	}
	root.Left = invertTree(root.Right)
	root.Right = invertTree(root.Left)
	return root
}

// Second idea
func invertTree(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	p := root
	for len(stack) > 0 || p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		p.Left, p.Right = p.Right, p.Left
		p = p.Left
	}
	return root
}
