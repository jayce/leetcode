/*
 * Creation in 2020-10-27 by Chuan Liu (Jayce)
 */

package leetcode

// RemoveNthFromEnd First idea
func RemoveNthFromEnd(head *ListNode, n int) *ListNode { // time: O(N)
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for i := 1; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// Second idea recursive
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dfs func(head *ListNode) *ListNode
	dfs = func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}

		head.Next = dfs(head.Next) // time: O(N), space: O(N)
		if n--; n != 0 {
			return head
		}
		return head.Next
	}

	return dfs(head)
}
