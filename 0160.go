/*
 * Creation in 2020-11-06 by Chuan Liu (Jayce)
 */

package leetcode

// GetIntersectionNode First idea
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	hash := make(map[*ListNode]bool)       // space: O(N)
	for p := headA; p != nil; p = p.Next { // time: O(N)
		hash[p] = true
	}

	for p := headB; p != nil; p = p.Next { // time: O(N)
		if hash[p] {
			return p
		}
	}
	return nil
}

// Second idea
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	alen, blen := listCount(headA), listCount(headB) // time: O(N+M)
	if alen == 0 || blen == 0 {
		return nil
	}

	if alen < blen {
		alen, blen = blen, alen
		headA, headB = headB, headA
	}

	for i := alen - blen; i > 0; i-- { // time: O(K)
		headA = headA.Next
	}

	for headA != headB { // time: O(max(N, M)-k)
		headA, headB = headA.Next, headB.Next
	}
	return headA
}

func listCount(head *ListNode) int {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	return n
}

//   | -- 4 -- | - 2 - |      round1  round2  round3  round4
// A: 1, 3, 5, 7, 9, 11    => (4) 9   (2) nil (2) 9   (2) nil
// B: 2, 4, 9, 11 | 1, 3 | => (4) nil (2) 5   (2) 9   (2) nil
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != pb { // time: O(N+M)
		pa = swapNext(pa, headB)
		pb = swapNext(pb, headA)
	}
	return pa
}

func swapNext(A, B *ListNode) *ListNode {
	if A != nil {
		return A.Next
	}
	return B
}
