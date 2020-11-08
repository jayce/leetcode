/*
 * Creation in 2020-11-08 by Chuan Liu (Jayce)
 */

package leetcode

// flatten First idea
func flatten(root *Node) *Node {
	for p := root; p != nil; { // time: O(N)
		next := p.Next
		if p.Child != nil {
			p.Next = flattn(p.Child) // time: O(M)
			p.Next.Prev = p
			p.Child = nil

			for p.Next != nil { // time: O(M)
				p = p.Next
			}

			if next != nil {
				p.Next = next
				p.Next.Prev = p
			}
		}
		p = next
	}
	return root
}

// Second idea optimize for the time complexity
func flatten(root *Node) *Node {
	flattenHelp(root) // time: O(N)
	return root
}

func flattenHelp(root *Node) (*Node, *Node) {
	dummy := &Node{Next: root}
	for p := root; p != nil; p = p.Next { // time: O(N+CN+CN..)
		chead, ctail := flattenHelp(p.Child)
		p.Child = nil
		if p.Next != nil && ctail != nil {
			p.Next.Prev = ctail
			ctail.Next = p.Next
		}

		if chead != nil {
			chead.Prev = p
			p.Next = chead
		}

		if ctail != nil {
			p = ctail
		}

		dummy.Next = p
	}
	return root, dummy.Next
}

// Third idea
func flatten(root *Node) *Node { // time: O(N)
	for p := root; p != nil; {
		if p.Child == nil {
			p = p.Next
			continue
		}

		cp := p.Child
		for cp.Next != nil { // time: O(CN)
			cp = cp.Next
		}

		cp.Next = p.Next
		if p.Next != nil {
			p.Next.Prev = cp
		}

		p.Next = p.Child // make time: O(CN) again
		p.Child.Prev = p
		p.Child = nil
	}

	return root
}
