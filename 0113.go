/*
 * Creation in 2020-11-10 by Chuan Liu (Jayce)
 */

package leetcode

// pathSum First idea
func pathSum(root *TreeNode, sum int) [][]int {
	values := make([][]int, 0)
	stack := make([]*TreeNode, 0) // space: O(H)
	vstack := make([]int, 0)      // space: O(H)

	last := root
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			vstack = append(vstack, root.Val)
			stack = append(stack, root)
			sum -= root.Val
			root = root.Left
		}

		root = stack[len(stack)-1]
		if sum == 0 && root.Right == nil && root.Left == nil {
			// copy vstack, time: O(H)
			values = append(values, append([]int{}, vstack...))
		}

		// to avoid loop the right again
		if root.Right != nil && root.Right != last {
			root = root.Right
			continue
		}

		stack = stack[:len(stack)-1] // pop
		vstack = vstack[:len(vstack)-1]
		sum += root.Val // recover
		last = root
		root = nil
	}
	return values
}

// Second idea
func pathSum(root *TreeNode, sum int) [][]int {
	values := make([][]int, 0)
	vstack := make([]int, 0) // space: O(H)

	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) { // time: O(N)
		if root == nil {
			return
		}

		vstack = append(vstack, root.Val)
		if sum == root.Val && root.Left == nil && root.Right == nil {
			values = append(values, append([]int{}, vstack...)) // time: O(H)
			vstack = vstack[:len(vstack)-1]                     // remove the last one
			return
		}

		dfs(root.Left, sum-root.Val)
		dfs(root.Right, sum-root.Val)
		vstack = vstack[:len(vstack)-1] // remove  the root
	}

	dfs(root, sum)
	return values
}
