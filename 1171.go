/*
 * Creation in 2020-10-26 by Chuan Liu (Jayce)
 */

package leetcode

// RemoveZeroSumSublists First idea
func RemoveZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	sp := dummy
	for sp != nil && sp.Next != nil { // time: O(N)
		sum := 0
		for p := sp.Next; p != nil; p = p.Next { // time: O(M) or O(n-(p-s))
			sum += p.Val
			if sum == 0 {
				sp.Next = p.Next
				break
			}
		}
		if sum != 0 {
			sp = sp.Next
		}
	}
	return dummy.Next
}

// Second
func removeZeroSumSublists(head *ListNode) *ListNode { // time: O(N*M)
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next
	ssum, fsum := 0, 0
	for slow != nil { // time: O(N)
		for ; fast != nil; fast = fast.Next { // time: O(M)
			fsum += fast.Val
			if fsum == ssum {
				slow.Next = slow.Next // why don't we break here?
			}
		}

		if slow = slow.Next; slow != nil { // change the pontier f
			ssum += slow.Val
			fsum = ssum
			fast = slow.Next
		}
	}

	return dummy.Next
}

// optimized the time complexity
func removeZeroSumSublists(head *ListNode) *ListNode { // time: O(2N)
	prefixSum := make(map[int]*ListNode) // space: O(N)
	dummy := &ListNode{Next: head}

	sum := 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		prefixSum[sum] = p
	}

	sum = 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		p.Next = prefixSum[sum].Next
	}
	return dummy.Next
}
