/*
 * Creation in 2020-10-16 by Chuan Liu (Jayce)
 */

package leetcode

// SortedListToBST First idea
func SortedListToBST(head *ListNode) *TreeNode {
	nums := make([]int, 0)                // space: O(N)
	for p := head; p != nil; p = p.Next { // time: O(N)
		nums = append(nums, p.Val)
	}

	//   0   1  2  3  4
	// -10, -3, 0, 5, 7
	return createBST(nums) // time: O(N)
}

func createBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	return &TreeNode{
		Val:   nums[middle],
		Left:  createBST(nums[0:middle]),
		Right: createBST(nums[middle+1:]),
	}
}

// SortedListToBSTII optimized for space
func SortedListToBSTII(head *ListNode) *TreeNode {
	return createBSTII(&TreeNode{Next: head}, 0, countList(head)-1)
}

func countList(head *ListNode) int { // time: O(N)
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	return n
}

func createBSTII(dummy *ListNode, si, ei int) *TreeNode { // time: O(N)
	if si > ei {
		return nil
	}

	mi := si + (ei-si)/2
	left := createBST(dummy, si, mi-1)
	rootVal := dummy.Next.Val // why need dummy?
	dummy.Next = dummy.Next.Next

	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: createBST(dummy, mi+1, ei),
	}
}
