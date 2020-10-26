/*
 * Creation in 2020-10-26 by Chuan Liu (Jayce)
 */

package leetcode

// DeleteDuplicates First idea
func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy

	// example: 1
	// d, 1, 1, 1, 2, 3
	// s        f
	//          s  f
	// example: 2
	// d, 1, 2, 3
	// s  f
	//    s  f
	for slow != nil {
		fast := slow.Next
		for fast != nil && fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}

		if slow.Next == fast {
			slow = slow.Next
		} else {
			slow.Next = fast.Next
		}
	}

	return dummy.Next
}

// Second idea
func deleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicatesHelp(head)
}

// 1, 1, 1, nil
//    h  n
// h
func deleteDuplicatesHelp(head *ListNode) *ListNode { // time: O(N)
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next != nil && head.Val == head.Next.Val {
		head.Next = deleteDuplicatesHelp(head.Next)
		if head.Next != nil && head.Val == head.Next.Val { // why need this?
			return head.Next.Next
		}
		return head.Next
	}
	head.Next = deleteDuplicatesHelp(head.Next)
	return head
}
