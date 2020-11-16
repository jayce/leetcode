/*
 * Creation in 2020-11-16 by Chuan Liu (Jayce)
 */

package leetcode

// isValidBST First idea
func isValidBST(root *TreeNode) bool {
	f := func(v int) bool {
		return true
	}
	return isValidBSTHelp(f, root) // time: O(N*N*N)
}

func isValidBSTHelp(f func(v int) bool, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !f(root.Val) {
		return false
	}
	fleft := func(v int) bool {
		return v < root.Val && f(v) // call times: O(N)
	}
	fright := func(v int) bool {
		return root.Val < v && f(v) // call times: O(N)
	}
	if root.Left != nil && !(root.Left.Val < root.Val) {
		return false
	}
	if root.Right != nil && !(root.Val < root.Right.Val) {
		return false
	}
	return isValidBSTHelp(fleft, root.Left) && // time: O(N)
		isValidBSTHelp(fright, root.Right)
}

// Second idea
func isValidBST(root *TreeNode) bool {
	min := ^int(^uint(0) >> 1)        // init: minimum
	return isValidBSTHelp(root, &min) // time: O(N)
}

func isValidBSTHelp(root *TreeNode, min *int) bool {
	if root == nil {
		return true
	}
	if !isValidBSTHelp(root.Left, min) || root.Val <= *min {
		return false
	}
	*min = root.Val
	return isValidBSTHelp(root.Right, min)
}

// Third idea
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	last := (*TreeNode)(nil)
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		if last != nil && last.Val >= root.Val {
			return false
		}
		stack = stack[:len(stack)-1]
		last = root
		root = root.Right
	}
	return true
}

// min < root < max
func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return dfs(root.Left, min, root) &&
		dfs(root.Right, root, max)
}
