/*
 * Creation in 2020-10-27 by Chuan Liu (Jayce)
 */

package leetcode

// DeleteNAfterM First idea
func DeleteNAfterM(head *ListNode, m, n int) *ListNode { // time: O(N)
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < n && fast.Next != nil; i++ {
		fast = fast.Next
	}

	for i := 0; i < m && slow.Next != nil; i++ {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	slow.Next = fast.Next
	return dummy.Next
}
