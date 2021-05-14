/*
 * Creation in 2021-05-13 by Chuan Liu (Jayce)
 */

package leetcode

// First idea
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{Next: list1}
	aprev := dummy
	index := 0
	for aprev.Next != nil && index < a { // time: O(a)
		aprev = aprev.Next
		index++
	}

	bnode := aprev.Next
	for bnode != nil && index < b { // time: O(b-a)
		bnode = bnode.Next
		index++
	}

	aprev.Next = list2

	if bnode == nil {
		return dummy.Next
	}

	for list2 != nil && list2.Next != nil { // time: O(N)
		list2 = list2.Next
	}

	list2.Next = bnode.Next
	return dummy.Next
}
