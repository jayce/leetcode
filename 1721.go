/*
 * Creation in 2021-05-14 by Chuan Liu (Jayce)
 */

package leetcode

// First idea Swapp Values
func swapNodes(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	first, slow, fast := dummy, dummy, dummy

	for i := 0; i < k; i++ { // time: O(k)
		fast = fast.Next
	}

	first = fast

	for fast != nil { // time: O(n-k)
		slow = slow.Next
		fast = fast.Next
	}

	first.Val, slow.Val = slow.Val, first.Val
	return head
}

// Second idea Swapp Nodes
func swapNodes(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	fprev, sprev, fast := dummy, dummy, dummy

	// start 1 to find the kth's prev
	for i := 0; i < k-1; i++ { // time: O(k)
		fast = fast.Next
	}

	fprev = fast
	fast = fast.Next

	for fast.Next != nil { // time: O(n-k)
		sprev = sprev.Next
		fast = fast.Next
	}

	// same node
	if fprev == sprev {
		return dummy.Next
	}

	first := fprev.Next
	second := sprev.Next

	//           (fprev)
	// 1. sprev->second->first->next
	if second == fprev {
		first, second = second, first
		fprev, sprev = sprev, fprev
	}

	//           (sprev)
	// 2. fprev->first->second->next
	if first == sprev {
		first.Next = second.Next
		second.Next = first
		fprev.Next = second
		return dummy.Next
	}

	// 3. fprev->first->..->second->next
	fnext, snext := first.Next, second.Next

	fprev.Next = second
	second.Next = fnext
	sprev.Next = first
	first.Next = snext
	return dummy.Next
}
