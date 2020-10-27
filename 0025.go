/*
 * Creation in 2020-10-27 by Chuan Liu (Jayce)
 */

package leetcode

// ReverseKGroup First idea
func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	for k > 1 { // time: O(n/k)
		for i := 0; i < k; i++ { // time: O(K)
			if fast == nil || fast.Next == nil {
				return dummy.Next
			}
			fast = fast.Next
		}

		B := fast.Next
		fast.Next = nil

		newHead := reverseList(slow.Next) // time: O(K)
		fast = slow.Next
		fast.Next = B
		slow.Next = newHead
		slow = fast
	}

	return dummy.Next
}

// Second idea recursive
func reverseKGroup(head *ListNode, k int) *ListNode { // time: O(2K * (n/k))
	if ok, nextHead := splitList(head, k); ok { // time: O(K)
		newHead := reverseList(head)           // time: O(K)
		head.Next = reverseKGroup(nextHead, k) // time: O(n/k)
		return newHead
	}
	return head
}

func splitList(head *ListNode, k int) (bool, *ListNode) {
	for i := 1; i < k; i++ {
		if head == nil {
			return false, nil
		}
		head = head.Next
	}
	next := head.Next
	head.Next = nil
	return true, next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

// Third idea
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	dhead := dummy
	for k > 1 { // time: O(n/k)
		if ok, newHead, nextHead := splitReverseList(head, k); !ok { // time: O(K)
			dhead.Next = reverseList(newHead) // time: O(n%k)
			break
		}
		dhead.Next = newHead // d->head
		dhead = head         // tail
		head = nextHead
	}

	return dummy.Next
}

func splitReverseList(head *ListNode, k int) (bool, *ListNode, *ListNode) {
	prev := (*ListNode)(nil)
	next := prev

	for i := 1; i < k; i++ {
		if head == nil {
			return false, nil
		}
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return true, prev, next
}
