/*
 * Creation in 2020-11-08 by Chuan Liu (Jayce)
 */

package leetcode

// splitListToParts First idea
func splitListToParts(root *ListNode, k int) []*ListNode {
	n := 0
	for p := root; p != nil; p = p.Next { // time: O(N)
		n++
	}

	results := make([]*ListNode, k)
	kn, remain := n/k, n%k
	for i, p := 0, root; p != nil; i++ { // time: O(N)
		results[i] = p
		jump := kn

		if remainder > 0 {
			remainder--
			jump++
		}

		// Note: split p and p.Next
		for jump -= 1; jump > 0 && p != nil; jump-- {
			p = p.Next
		}

		next := p.Next
		p.Next = nil
		p = next
	}

	return results
}

// Second idea optimize for time: O(N+K)
func splitListToParts(root *ListNode, k int) []*ListNode {
	nodes := make([]*ListNode, 0, 10)     // space: O(N)
	for p := root; p != nil; p = p.Next { // time: O(N)
		nodes = append(nodes, p)
	}

	results := make([]*ListNode, k)
	kn, remainder := len(nodes)/k, len(nodes)%k
	for i := 0; len(nodes) > 0; i++ { // time: O(K)
		results[i] = nodes[0]

		jump := kn
		if remainder > 0 {
			remainder--
			jump++
		}

		nodes[jump-1].Next = nil // split
		nodes = nodes[jump:]
	}
	return results
}
