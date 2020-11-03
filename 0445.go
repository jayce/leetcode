/*
 * Creation in 2020-11-03 by Chuan Liu (Jayce)
 */

package leetcode

// AddTwoNumbers First idea
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := toSlice(l1) // time: O(N), space: O(N)
	s2 := toSlice(l2) // time: O(M), space: O(M)

	var l3 *ListNode
	p1, p2 := len(s1), len(s2)
	remain := 0
	for p1 > 0 || p2 > 0 || remain > 0 { // time: O(max(N, M))
		val := 0
		if p1 > 0 {
			val += s1[p1-1].Val
			p1--
		}
		if p2 > 0 {
			val += s2[p2-1].Val
			p2--
		}

		val += remain
		remain = val / 10
		val %= 10

		l3 = &ListNode{Val: val, Next: l3}
	}
	return l3
}

func toSlice(head *ListNode) []*ListNode {
	s := make([]*ListNode, 0, 10)
	for ; head != nil; head = head.Next {
		s = append(s, head)
	}
	return s
}
