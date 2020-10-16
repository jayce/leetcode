/*
 * Creation in 2020-10-16 by Chuan Liu (Jayce)
 */

package leetcode

// OddEvenList First idea
func OddEvenList(head *ListNode) *ListNode {
	odd, even := &ListNode{Next: head}, &ListNode{Next: head}
	op, ep := odd, even
	for op.Next != nil && op.Next.Next != nil { // time: O(N)
		op.Next = ep.Next
		op = op.Next

		ep.Next = op.Next
		ep = ep.Next
	}
	op.Next = even.Next
	return odd.Next
}

// OddEvenListII recursive
func OddEvenListII(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	oddEnd, evenHead := oddEvenListHelp(head, head.Next) // time: O(N)
	oddEnd.Next = evenHead
	return head
}

func oddEvenListHelp(odd, even *ListNode) (*ListNode, *ListNode) {
	if odd == nil || even == nil || even.Next == nil {
		return odd, even
	}
	odd.Next = even.Next
	odd, even.Next = oddEvenListHelp(even.Next, even.Next.Next)
	return odd, even
}
