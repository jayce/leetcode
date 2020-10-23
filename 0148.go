/*
 * Creation in 2020-10-23 by Chuan Liu (Jayce)
 */

package leetcode

// SortList First idea for section
func SortList(head *ListNode) *ListNode { // time: O(N*N)
	dummy := &ListNode{}
	left, right := dummy, &ListNode{}

	dp := &ListNode{Next: head}
	for dp.Next != nil { // time: O(N)
		minPrev, maxPrev := dp, dp
		for p := dp.Next; p.Next != nil; p = p.Next { // time: O(N) O(N-2) .. 0
			switch {
			case minPrev.Next.Val >= p.Next.Val:
				minPrev = p
			case maxPrev.Next.Val < p.Next.Val:
				maxPrev = p
			}
		}

		minNode := minPrev.Next
		minPrev.Next = minPrev.Next.Next
		left.Next = minNode
		left = minNode

		if minPrev == maxPrev {
			break
		}

		if minPrev.Next == maxPrev.Next { // why?
			// if we had d->2->3, the 3 will be handled two times
			// first 3 is the maximum, second it is the minimum
			continue
		}

		maxNode := maxPrev.Next
		maxPrev.Next = maxPrev.Next.Next
		maxNode.Next = right.Next
		right.Next = maxNode
	}

	left.Next = right.Next
	return dummy.Next
}

// SortListII mergeSort
func SortListII(head *ListNode) *ListNode { // time: O(N * logN)
	if head == nil || head.Next == nil {
		return head
	}

	s, f := head, head.Next
	for f != nil && f.Next != nil { // time: O(N/2)
		s = s.Next
		f = f.Next.Next
	}

	f = s.Next
	s.Next = nil
	dummy := &ListNode{}
	mergeList(dummy, SortListII(head), SortListII(f)) // time: O(N)
	return dummy.Next
}

func mergeList(dummy, l1, l2 *ListNode) {
	if l1 == nil {
		dummy.Next = l2
		return
	}

	if l2 == nil {
		dummy.Next = l1
		return
	}

	if l1.Val > l2.Val {
		dummy.Next = l1
		mergeList(dummy.Next, l1.Next, l2)
		return
	}
	dummy.Next = l2
	mergeList(dummy.Next, l1, l2.Next)
	return
}
