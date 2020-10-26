/*
 * Creation in 2020-10-26 by Chuan Liu (Jayce)
 */

package leetcode

// MiddleNode First idea
func MiddleNode(head *ListNode) *ListNode {
	n, node := 0, head
	for head != nil {
		n++
		head = head.Next
	}
	for i := 0; node != nil && i <= n/2; i++ {
		node = node.Next
	}
	return node
}

// Second idea
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil { // time: O(N/2)
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// recursive
func middleNode(head *ListNode) *ListNode {
	return recursive(head, head)
}

func recursive(slow, fast *ListNode) *ListNode {
	if fast == nil || fast.Next == nil {
		return slow
	}
	return recursive(slow.Next, fast.Next.Next)
}
