/*
 * Creation in 2021-05-17 by Chuan Liu (Jayce)
 */

package leetcode

// First idea
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	return BSTIterator{stack}
}

func (this *BSTIterator) Next() int {
	last := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	root := last.Right
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
	return last.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
