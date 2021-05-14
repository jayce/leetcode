/*
 * Creation in 2021-05-13 by Chuan Liu (Jayce)
 */

package leetcode

// First idea
type FrontMiddleBackQueue struct {
	dummy *ListNode
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{&ListNode{}}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.pushIndexAt(0, val)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	this.pushIndexAt(this.dummy.Val/2, val)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.pushIndexAt(this.dummy.Val, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	return this.popIndexAt(0)
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	return this.popIndexAt((this.dummy.Val - 1) / 2) // why -1?
}

func (this *FrontMiddleBackQueue) PopBack() int {
	return this.popIndexAt(this.dummy.Val - 1)
}

func (this *FrontMiddleBackQueue) pushIndexAt(n, val int) {
	dummy := this.findIndexPrev(n)
	dummy.Next = &ListNode{
		Next: dummy.Next,
		Val:  val,
	}

	this.dummy.Val++
}

func (this *FrontMiddleBackQueue) popIndexAt(n int) int {
	dummy := this.findIndexPrev(n)
	next := dummy.Next
	if next == nil {
		return -1
	}

	this.dummy.Val--
	dummy.Next = next.Next
	return next.Val
}

func (this *FrontMiddleBackQueue) findIndexPrev(i int) *ListNode {
	if i < 0 || i > this.dummy.Val {
		return &ListNode{}
	}

	p := this.dummy
	for ; p.Next != nil && i > 0; i-- {
		p = p.Next
	}
	return p
}
