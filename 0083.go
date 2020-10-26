/*
 * Creation in 2020-10-26 by Chuan Liu (Jayce)
 */

package leetcode

// DeleteDuplicates First idea
func DeleteDuplicates(head *ListNode) *ListNode {
	for p := head; p != nil && p.Next != nil; { // time: O(N)
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

// DeleteDuplicates Second idea for recursive
func deleteDuplicates(head *ListNode) *ListNode {
	deleteDuplicatesHelp(head) // time: O(N)
	return head
}

func deleteDuplicatesHelp(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
	} else {
		head = head.Next
	}
	deleteDuplicatesHelp(head)
	return
}
