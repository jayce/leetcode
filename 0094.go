/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0) // time: O(H)
	values := make([]int, 0)
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root) // push root
			root = root.Left            // left
		}
		root = stack[len(stack)-1]        // pop root
		stack = stack[:len(stack)-1]      // trim the last
		values = append(values, root.Val) // handle root
		root = root.Right                 // right
	}
	return values
}

// Second idea
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(inorderTraversal(root.Left),
		root.Val, inorderTraversal(root.Right)...)
}

// Third idea optimized malloc
func inorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	inorderTraversalHelp(root, &values)
	return values
}

func inorderTraversalHelp(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inorderTraversalHelp(root.Left, values)
	*values = append(*values, root.Val)
	inorderTraversalHelp(root.Right, values)
	return
}
