/*
 * Creation in 2020-11-16 by Chuan Liu (Jayce)
 */

package leetcode

// sortedArrayToBST First idea
func sortedArrayToBST(nums []int) *TreeNode { // time: O(logN)
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[0:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}

type item struct {
	node  *TreeNode
	start int
	end   int
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{}
	stack := []*item{}
	stack = append(stack, &item{root, 0, len(nums) - 1})
	for len(stack) > 0 {
		it := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		m := it.start + (it.end-it.start)/2
		it.node.Val = nums[m]
		if it.start < m {
			it.node.Left = &TreeNode{}
			stack = append(stack, &item{it.node.Left, it.start, m - 1})
		}
		if it.end > m {
			it.node.Right = &TreeNode{}
			stack = append(stack, &item{it.node.Right, m + 1, it.end})
		}
	}

	return root
}
