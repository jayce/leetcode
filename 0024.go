/*
 * Creation in 2020-10-28 by Chuan Liu (Jayce)
 */

package leetcode

// SwapPairs First idea top-bottom
func SwapPairs(head *ListNode) *ListNode { // time: O(N)
	dummy := &ListNode{Next: head}
	// a->b->c->d
	for p := dummy; head != nil && head.Next != nil; head = head.Next {
		p.Next = head.Next         // first: p->b->c->d
		head.Next = head.Next.Next // head: a->c->d
		p.Next.Next = head         // second: p->b->a->c->d
		p = p.Next.Next            // p: a->c->d
	}
	return dummy.Next
}

// Second idea bottom-top
func swapPairs(a *ListNode) *ListNode { // time: O(N/2)
	if a == nil || a.Next == nil {
		return a
	}
	b := a.Next
	a.Next = swapPairs(a.Next.Next)
	b.Next = a
	return b
}
