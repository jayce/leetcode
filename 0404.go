/*
 * Creation in 2020-11-17 by Chuan Liu (Jayce)
 */

package leetcode

// sumOfLeftLeaves First idea
func sumOfLeftLeaves(root *TreeNode) int { // time: O(N)
	if root == nil {
		return 0
	}

	if root.Left != nil &&
		root.Left.Left == nil &&
		root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

// Second idea another version from the First
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			if root.Left != nil &&
				root.Left.Left == nil &&
				root.Left.Right == nil {
				sum += root.Left.Val
			}
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return sum
}
