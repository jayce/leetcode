/*
 * Creation in 2020-10-15 by Chuan Liu (Jayce)
 */

package leetcode

type MyLinkedList struct {
	length int
	dummy  *ListNode
}

/* Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{dummy: &ListNode{}}
}

// find the node at the index-1, but the p.Next is the real node.
// complexity: O(N)
func (this *MyLinkedList) findAtIndex(index int) *ListNode {
	if index < 0 || index > this.length {
		return nil
	}

	p := this.dummy
	for index > 0 {
		index--
		p = p.Next
	}
	return p
}

/* Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
// complexity: O(N)
func (this *MyLinkedList) Get(index int) int {
	p := this.findAtIndex(index)
	if p == nil || p.Next == nil {
		return -1
	}
	return p.Next.Val
}

/* Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
// complexity: O(1)
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
	return
}

/* Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.length, val)
	return
}

/* Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	p := this.findAtIndex(index)
	if p == nil {
		return
	}
	this.length++
	p.Next = &ListNode{Val: val, Next: p.Next}
	return
}

// DeleteAtIndex Delete the index-th node in the linked list, if the index is valid.
func (this *MyLinkedList) DeleteAtIndex(index int) {
	p := this.findAtIndex(index)
	if p == nil || p.Next == nil {
		return
	}
	this.length--
	p.Next = p.Next.Next
	return
}
