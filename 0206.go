/*
 * Creation in 2020-10-23 by Chuan Liu (Jayce)
 */

package leetcode

// ReverseList First idea
func ReverseList(head *ListNode) *ListNode { // time: O(N)
	var previous *ListNode
	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}
	return previous
}

// ReverseListII reverse version
func ReverseListII(head *ListNode) *ListNode { // time: O(N)
	if head == nil || head.Next == nil {
		return head
	}
	node := ReverseListII(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

// ReverseListIII optimized for ReverseListII
// example: head = ReverseListIII(nil, head)
func ReverseListIII(A, B *ListNode) *ListNode { // time: O(N)
	if B == nil {
		return A
	}
	next := B.Next
	B.Next = A
	return ReverseListIII(B, next)
}
