/*
 * Creation in 2020-10-23 by Chuan Liu (Jayce)
 */

package leetcode

// Partition First idea
func Partition(head *ListNode, x int) *ListNode {
	dummyMin, dummyMax := &ListNode{}, &ListNode{}
	minp, maxp := dummyMin, dummyMax
	for p := head; p != nil; p = p.Next { // time: O(N)
		if p.Val < x {
			minp.Next = p
			minp = p
		} else {
			maxp.Next = p
			maxp = p
		}
	}

	maxp.Next = nil // to avoid cycle of the list
	minp.Next = dummyMax.Next
	return dummyMin.Next
}
