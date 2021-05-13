/*
 * Creation in 2021-05-13 by Chuan Liu (Jayce)
 */

package leetcode

// First idea space: O(1)
func detectCycle(head *ListNode) *ListNode {
	// 1, 2, 3, 4, 5 -> 3
	// - k -   m   -
	// n = k + m
	slow, fast := head, head
	for fast != nil && fast.Next != nil { // time: O((N+M)/2)
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	cycleLen := 1                               // slow is the first
	for p := slow; p.Next != slow; p = p.Next { // time: O(M-1)
		cycleLen++
	}

	slow, fast = head, head
	for ; cycleLen > 0; cycleLen-- { // time: O(M)
		fast = fast.Next
	}

	for slow != fast { // time: O(N-M)
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
