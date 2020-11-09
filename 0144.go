/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

// preorderTraversal First idea
func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0) // space: O(H)
	values := make([]int, 0)
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			values = append(values, root.Val) // handle root
			stack = append(stack, root)       // push root
			root = root.Left                  // left
		}

		root = stack[len(stack)-1]   // root
		stack = stack[:len(stack)-1] // pop
		root = root.Right            // right
	}
	return values
}

// Second idea recursive
func preorderTraversal(root *TreeNode) []int { // time: O(N)
	if root == nil {
		return []int{}
	}

	values := []int{root.Val}
	values = append(values, preorderTraversal(root.Left)...)
	return append(values, preorderTraversal(root.Right)...)
}
