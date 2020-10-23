/*
 * Creation in 2020-10-23 by Chuan Liu (Jayce)
 */

package leetcode

// ReorderList First Idea
func ReorderList(head *ListNode) {
	elmts := []*ListNode{}
	for p := head; p != nil; p = p.Next {
		elmts = append(elmts, p)
	}

	n := len(elmts) - 1
	p := head
	for p.Next != nil && p.Next.Next != nil {
		elmts[n].Next = p.Next
		p.Next = elmts[n]

		n--
		elmts[n].Next = nil
		p = p.Next.Next
	}

	// another merge code
	// for h, t := 0, len(elmts)-1; h < t; h, t = h+1, t-1 {
	// 	elmts[h].Next = elmts[t]
	// 	elmts[t].Next = elmts[h+1]
	// 	elmts[h+1].Next = nil
	// }
}

func ReorderListII(head *ListNode) {
	if head == nil {
		return
	}
	s, f, head := head, head, nil
	p := head // nil
	for f != nil && f.Next != nil {
		f = f.Next.Next
		p, s = popInsertFromList(p, s)
	}
	if f != nil {
		head, s = popInsertFromList(head, s)
	}
	for p != nil {
		head, s = popInsertFromList(head, s)
		head, p = popInsertFromList(head, p)
	}
}

func ReorderListII(head *ListNode) {
	reorderListHelp(nil, head, head)
}

func reorderListHelp(p, s, f *ListNode) (*ListNode, *ListNode) {
	if f == nil || f.Next == nil {
		if f != nil {
			p, s = popInsertFromList(p, s)
		}
		return p, s
	}
	m := s
	p, m = reorderListHelp(p, s.Next, f.Next.Next)
	p, m = popInsertFromList(p, m)
	p, s = popInsertFromList(p, s)
	return p, m
}

func popInsertFromList(h, p *ListNode) (*ListNode, *ListNode) {
	if p == nil {
		return h, nil
	}
	next := p.Next
	p.Next = h
	return p, next
}
