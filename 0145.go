/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

// postorderTraversal First idea
func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0) // space: O(H)
	values := make([]int, 0)
	for len(stack) > 0 || root != nil { // time: O(N) relloc: N times
		for root != nil {
			values = append([]int{root.Val}, values...) // handle root
			stack = append(stack, root)                 // push root
			root = root.Right                           // right
		}

		root = stack[len(stack)-1]   // root
		stack = stack[:len(stack)-1] // pop
		root = root.Left             // left
	}
	return values
}

// Second optimized malloc
func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0) // space: O(H)
	values := make([]int, 0)
	last := (*TreeNode)(nil)
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root) // push root
			root = root.Left
		}

		root = stack[len(stack)-1]
		if root.Right != nil && root.Right != last { // not the right node
			root = root.Right
			continue
		}

		stack = stack[:len(stack)-1] // pop
		values = append(values, root.Val)
		last = root
		root = nil // avoid loop
	}
	return values
}

// recursive
func postorderTraversal(root *TreeNode) []int { // time: O(N)
	if root == nil {
		return []int{}
	}
	return append(append(postorderTraversal(root.Left),
		postorderTraversal(root.Right)...), root.Val)
}
