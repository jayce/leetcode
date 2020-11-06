/*
 * Creation in 2020-11-06 by Chuan Liu (Jayce)
 */

package leetcode

// RemoveElements First idea
func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for p := dummy; p != nil; p = p.Next { // time: O(N)
		for p.Next != nil && p.Next.Val == val {
			p.Next = p.Next.Next
		}
	}
	return dummy.Next
}

// Second idea
func removeElements(head *ListNode, val int) *ListNode { // time: O(N)
	if head == nil {
		return nil
	}

	if head.Val == val {
		return removeElements(head.Next, val) // space: O(N)
	}
	head.Next = removeElements(head.Next, val)
	return head
}
