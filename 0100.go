/*
 * Creation in 2020-11-15 by Chuan Liu (Jayce)
 */

package leetcode

// isSameTree First idea
func isSameTree(p *TreeNode, q *TreeNode) bool { // time: O(N)
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

// Second idea
func isSameTree(r1 *TreeNode, r2 *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || r1 != nil || r2 != nil { // time: O(N)
		for r1 != nil || r2 != nil {
			if r1 == nil || r2 == nil || r1.Val != r2.Val {
				return r1 == r2
			}
			stack = append(stack, r1, r2)
			r1, r2 = r1.Left, r2.Left
		}

		r2 = stack[len(stack)-1]
		r1 = stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		r1, r2 = r1.Right, r2.Right
	}
	return true
}
