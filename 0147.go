/*
 * Creation in 2020-10-26 by Chuan Liu (Jayce)
 */

package leetcode

// InsertionSortList Fist idea
func InsertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for head != nil { // time: O(N)
		next := head.Next
		if p.Val >= head.Val {
			p = dummy
		}

		for p.Next != nil && p.Next.Val < head.Val {
			p = p.Next
		}

		head.Next = p.Next
		p.Next = head
		head = next
	}

	return dummy.Next
}

// InsertionSortListII Second idea
func InsertionSortListII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	d := &ListNode{Next: InsertionSortListII(head.Next)}
	p := d
	for p.Next != nil && p.Next.Val < head.Val {
		p = p.Next
	}
	head.Next = p.Next
	p.Next = head
	return d.Next
}

// InsertionSortListIII recursive
func InsertionSortListIII(head *ListNode) *ListNode {
	dummy := &ListNode{}
	return insertionSortListHelp(dummy, head)
}

func insertionSortListHelp(d, head *ListNode) *ListNode {
	if head == nil {
		return d.Next
	}
	p := d
	for p.Next != nil && p.Next.Val < head.Val {
		p = p.Next
	}
	next := head.Next
	head.Next = p.Next
	p.Next = head
	return insertionSortListHelp(p, next)
}
