/*
 * Creation in 2020-11-15 by Chuan Liu (Jayce)
 */

package leetcode

// findSecondMinimumValue First idea
func findSecondMinimumValue(root *TreeNode) int {
	maxInt := int(^uint(0) >> 1) // init: minimum
	first, second := root.Val, maxInt
	var f func(root *TreeNode) // time: O(N)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		if first < root.Val && root.Val < second {
			second = root.Val
			return
		}
		f(root.Left)
		f(root.Right)
	}
	f(root)
	if second < maxInt {
		return second
	}
	return -1
}

// Second idea
func findSecondMinimumValue(root *TreeNode) int {
	first := root.Val
	var f func(root *TreeNode)
	f = func(root *TreeNode) int { // time: O(N)
		if root == nil {
			return -1
		}
		if first < root.Val {
			return root.Val
		}
		L := f(root.Left)
		R := f(root.Right)
		if L == -1 || R == -1 {
			return maxmin(L, R, L > R)
		}
		return maxmin(L, R, L < R)
	}
	return f(root)
}

func maxmin(a, b int, x bool) int {
	if x {
		return a
	}
	return b
}
