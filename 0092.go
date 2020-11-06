/*
 * Creation in 2020-11-06 by Chuan Liu (Jayce)
 */

package leetcode

// reverseBetween First idea
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Next: head}
	dp := dummy

	for i := 0; i < m; i++ { // time: O(M)
		dp = dp.Next
	}

	// reverse
	var rh, next *ListNode
	rp := dp.Next
	for s := m; rp != nil && s < n; s++ { // time: O(N)
		next = rp.Next
		rp.Next = rh
		rh = rp
		rp = next
	}

	if next != nil && dp.Next != nil {
		dp.Next.Next = next
	}

	dp.Next = rh
	return dummy.Next
}

// Second idea from reverseI
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		var reverseN func(head *ListNode, n int) *ListNode
		var lastHead *ListNode
		reverseN = func(head *ListNode, n int) *ListNode {
			if n == 1 {
				lastHead = head.Next
				return head
			}
			newHead := reverseN(head.Next, n-1)
			head.Next.Next = head
			head.Next = lastHead
			return newHead
		}
		return reverseN(head, n)
	}

	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}
