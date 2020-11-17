/*
 * Creation in 2020-11-17 by Chuan Liu (Jayce)
 */

package leetcode

// constructMaximumBinaryTree First idea
func constructMaximumBinaryTree(nums []int) *TreeNode { // time: O(N*N)
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	m := 0
	for i := 1; i < len(nums); i++ { // time: O(N)
		if nums[m] < nums[i] {
			m = i
		}
	}

	return &TreeNode{
		Val:   nums[m],
		Left:  constructMaximumBinaryTree(nums[0:m]),  // time: O(N/2)
		Right: constructMaximumBinaryTree(nums[m+1:]), // time: O(N/2)
	}
}
