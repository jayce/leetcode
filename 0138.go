/*
 * Creation in 2020-10-15 by Chuan Liu (Jayce)
 */

package leetcode

// CopyRandomList First idea
func CopyRandomList(head *Node) *Node {
	mapToIndex := make(map[*Node]int) // space: O(N)
	IndexToNode := make([]*Node, 0)   // space: O(N)
	index := 0
	dummy := &Node{}
	dp := dummy
	for p := head; p != nil; p = p.Next { // time: O(N)
		mapToIndex[p] = index
		index++
		node := *p
		IndexToNode = append(IndexToNode, &node)
		dp.Next = &node
		dp = dp.Next
	}

	for p := dummy.Next; p != nil; p = p.Next { // time: O(N)
		if ok, index := mapToIndex[p.Random]; ok {
			p.Random = IndexToNode[index]
		}
	}
	return dummy.Next
}

// CopyRandomListII optimized for Space
func CopyRandomListII(head *Node) *Node {
	if head == nil {
		return head
	}
	for p := head; p != nil; p = p.Next.Next { // time: O(N)
		copy := *p
		p.Next = &copy
	}
	for p := head; p != nil; p = p.Next.Next { // time: O(N)
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}
	// a1-a2-b1-b2
	//  p  pc
	copy := head.Next
	p, pc := head, copy
	for p != nil && p.Next != nil { // time: O(N)
		p.Next = pc.Next   // a1-b1
		p, pc = pc, p.Next // a2, b1
	}
	return copy
}
