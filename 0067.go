/*
 * Creation in 2020-11-05 by Chuan Liu (Jayce)
 */
package leetcode

// RotateRight First idea
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverse(head) // time: O(N)

	head.Next = newHead
	for ; k > 1; k-- { // time: O(K)
		newHead = newHead.Next
	}

	head = newHead.Next
	newHead.Next = nil
	return reverse(head) // time: O(N)
}

func reverse(head *ListNode) *ListNode {
	var rh *ListNode
	for head != nil {
		next := head.Next
		head.Next = rh
		rh = head
		head = next
	}
	return rh
}

// Second idea
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	tail, n := head, 0
	for tail != nil { // time: O(N)
		n++
		tail = tail.Next
	}

	tail.Next = head

	for steps := n - k%n; steps > 0; steps-- { // time: O(N-K%N)
		tail = tail.Next
	}

	head, tail.Next = tail.Next, nil
	return head
}
