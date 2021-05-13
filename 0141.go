/*
 * Creation in 2020-11-05 by Chuan Liu (Jayce)
 */

package leetcode

// First idea
func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]bool)      // space: O(N-k/2)
	for ; head != nil; head = head.Next { // time: O(N)
		if hash[head] {
			return true
		}
	}
	return false
}

// Second idea
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { // time: O(N)
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Third idea
func hasCycle(head *ListNode) bool { // time: O(N)
	dummy := &ListNode{Next: head}
	return hasCycleHelp(dummy, dummy.Next)
}

func hasCycleHelp(slow, fast *ListNode) bool {
	if fast != nil && fast.Next != nil {
		return slow == fast ||
			hasCycleHelp(slow.Next, fast.Next.Next)
	}
	return slow == fast
}
