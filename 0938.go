/*
 * Creation in 2020-11-17 by Chuan Liu (Jayce)
 */

package leetcode

// rangeSum First idea
func rangeSumBST(root *TreeNode, L int, R int) int { // time: O(N)
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}

	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}

	return root.Val +
		rangeSumBST(root.Left, L, R) +
		rangeSumBST(root.Right, L, R)
}

// Second idea top-bottom
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if L <= root.Val && root.Val <= R {
		sum = root.Val
	}
	if L <= root.Val {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Val <= R {
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}
