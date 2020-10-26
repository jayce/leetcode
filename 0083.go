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
	}
	deleteDuplicatesHelp(head.Next)
	return
}

func deleteDuplicatesHelpII(head *ListNode) *ListNode { // time: O(N)
	if head == nil || head.Next == nil {
		return head
	}

	if head.Next != nil && head.Val == head.Next.Val {
		return deleteDuplicatesHelp(head.Next)
	}
	head.Next = deleteDuplicatesHelp(head.Next)
	return head
}
