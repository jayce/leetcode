/*
 * Creation in 2020-10-27 by Chuan Liu (Jayce)
 */

package leetcode

// MergeTwoLists First idea
func MergeTwoLists(A, B *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy

loop:
	for { // time: O(N)
		switch {
		case A == nil:
			p.Next = B
			break loop
		case B == nil:
			p.Next = A
			break loop
		case A.Val < B.Val:
			p.Next = A
			p, A = A, A.Next
		case A.Val >= B.Val:
			p.Next = B
			p, B = B, B.Next
		}
	}

	return dummy.Next
}

// Second idea recursive
func mergeTwoLists(A, B *ListNode) *ListNode { // time: O(N)
	if A == nil {
		return B
	}
	if B == nil {
		return A
	}

	if A.Val > B.Val {
		B.Next = mergeTwoLists(A, B.Next)
		return B
	}
	A.Next = mergeTwoLists(A.Next, B)
	return A
}

// Third idea recursive at the tail
func mergeTwoLists(A, B *ListNode) *ListNode {
	dummy := &ListNode{}
	mergeTwoListsII(dummy, A, B)
	return dummy.Next
}

func mergeTwoListsII(dummy, A, B *ListNode) *ListNode {
	if A == nil {
		dummy.Next = B
		return A
	}
	if B == nil {
		dummy.Next = A
		return B
	}
	if A.Val > B.Val {
		dummy.Next = B
		return mergeTwoListsII(dummy.Next, A, B.Next)
	}
	dummy.Next = A
	return mergeTwoListsII(dummy.Next, A.Next, B)
}
