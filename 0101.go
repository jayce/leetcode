/*
 * Creation in 2020-11-19 by Chuan Liu (Jayce)
 */

package leetcode

// isSymmetric First idea
func isSymmetric(root *TreeNode) bool {
	return root == nil || dfs(root.Left, root.Right)
}

func dfs(L, R *TreeNode) bool { // time: O(N)
	if L == nil || R == nil {
		return L == R
	}
	return L.Val == R.Val &&
		dfs(L.Right, R.Left) &&
		dfs(L.Left, R.Right)
}

// Second idea
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{} // space: O(H)
	L, R := root.Left, root.Right

	for len(stack) > 0 || L != nil || R != nil { // time: O(N)
		for L != nil || R != nil {
			if L == nil || R == nil {
				return L == R
			}
			if L.Val != R.Val {
				return false
			}
			stack = append(stack, L, R)
			L, R = L.Left, R.Right
		}

		R = stack[len(stack)-1]
		L = stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		L, R = L.Right, R.Left
	}
	return true
}
