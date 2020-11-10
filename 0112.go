/*
 * Creation in 2020-11-10 by Chuan Liu (Jayce)
 */

package leetcode

// hasPathSum First idea
func hasPathSum(root *TreeNode, sum int) bool { // time: O(N)
	if root == nil {
		return false
	}
	// the last one
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)
}

// Second idea iteratively
func hasPathSum(root *TreeNode, sum int) bool {
	stack := make([]*TreeNode, 0)       // space: O(H)
	last := root                        // just for init
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root)
			sum -= root.Val // minus
			root = root.Left
		}

		root = stack[len(stack)-1]
		if sum == 0 && root.Left == nil && root.Right == nil {
			return true
		}

		// to avoid handle the right again
		if root.Right != nil && root.Right != last {
			root = root.Right
			continue
		}

		stack = stack[:len(stack)-1] // pop
		sum += root.Val              // recover
		last = root
		root = nil
	}
	return false
}
