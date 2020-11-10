/*
 * Creation in 2020-11-10 by Chuan Liu (Jayce)
 */

package leetcode

// sumRootToLeaf First idea
func sumRootToLeaf(root *TreeNode) int {
	return sumRootToLeafHelp(root, 0) // time: O(N)
}

func sumRootToLeafHelp(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = sum<<1 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return sumRootToLeafHelp(root.Left, sum) + sumRootToLeafHelp(root.Right, sum)
}
