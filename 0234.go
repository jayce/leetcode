/*
 * Creation in 2020-11-05 by Chuan Liu (Jayce)
 */

package leetcode

// isPalindrome First idea
func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{Next: head}
	stack := make([]*ListNode, 0, 10) // space: O(N)
	slow, fast := dummy, dummy
	// linked: d, 1, 2, 3, 4
	// round 1:d  s  f
	// round 2:      s     f
	for fast != nil && fast.Next != nil { // time: O(N/2)
		fast = fast.Next.Next
		slow = slow.Next
		stack = append(stack, slow)
	}

	slow = slow.Next
	for i := len(stack) - 1; i > 0; i-- { // time: O(N/2)
		if stack[i].Val != slow.Val {
			return false
		}
		slow = slow.Next
	}
	return true
}

// Second idea
func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{Next: head}
	fast := dummy
	head = nil

	for fast != nil && fast.Next != nil { // time: O(N/2)
		fast = fast.Next.Next

		next := dummy.Next.Next
		dummy.Next.Next = head
		head = dummy.Next
		dummy.Next = next
	}

	return compare(head, dummy.Next) // time: O(N/2)
}

func compare(A, B *ListNode) bool {
	if A == nil || B == nil {
		return A == B
	}

	return A.Val == B.Val && compare(A.Next, B.Next)
}
