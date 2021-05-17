/*
 * Creation in 2020-11-23 by Chuan Liu (Jayce)
 */

package leetcode

// kthSmallest First idea
func kthSmallest(root *TreeNode, k int) int {
	var inorder func(root *TreeNode) int
	inorder = func(root *TreeNode) int { // time: O(N)
		if root == nil {
			return 0
		}
		kth := inorder(root.Left)
		k -= 1
		if k < 0 {
			return kth
		}
		if k == 0 {
			return root.Val
		}
		return inorder(root.Right)
	}
	return inorder(root)
}

// Second idea
func kthSmallest(root *TreeNode, k int) int {
	kthNode := &TreeNode{}
	stack := []*TreeNode{}              // space: O(H)
	for len(stack) > 0 || root != nil { // time: O(N)
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k -= 1; k == 0 {
			kthNode = root
			break
		}
		root = root.Right
	}
	return kthNode.Val
}

// Third idea from leetcode
func kthSmallest(dummy *TreeLinkedNode, k int) int {
	for p := dummy.Next; p != dummy; p = p.Next {
		if k -= 1; k == 0 {
			return p.Val
		}
	}
	return -1
}

// BST with Double Linked
type TreeLinkedNode struct {
	Val   int
	Left  *TreeLinkedNode
	Right *TreeLinkedNode
	Prev  *TreeLinkedNode
	Next  *TreeLinkedNode
}

func NewTreeLinkedNode() *TreeLinkedNode {
	// dummy.Left is the root of tree
	dummy := &TreeLinkedNode{}
	initLink(dummy)
	return dummy
}

func (dummy *TreeLinkedNode) Search(val int) *TreeLinkedNode {
	for p := dummy.Left; p != nil; {
		switch {
		case p.Val > val:
			p = p.Left
		case p.Val < val:
			p = p.Right
		default:
			return p
		}
	}
	return nil
}

func (dummy *TreeLinkedNode) Insert(val int) *TreeLinkedNode {
	p := dummy.Left
	newNode := &TreeLinkedNode{Val: val}

	if p == nil {
		dummy.Left = newNode
		insertLink(dummy, newNode)
		return newNode
	}

	for p != nil { // time: O(logN), O(N)
		if p.Val == val {
			return p
		}

		if p.Val > val {
			if p.Left != nil {
				p = p.Left
				continue
			}

			// p.Left == nil
			p.Left = newNode
			break

		} else {
			if p.Right != nil {
				p = p.Right
				continue
			}

			// p.Right == nil
			p.Right = newNode
			break
		}
	}

	insertLink(dummy, newNode) // time: O(N)
	return newNode
}

func (dummy *TreeLinkedNode) Delete(val int) *TreeLinkedNode {
	pp := &dummy.Left
	for *pp != nil && (*pp).Val != val {
		switch {
		case (*pp).Val > val:
			pp = &(*pp).Left
		case (*pp).Val < val:
			pp = &(*pp).Right
		}
	}

	if (*pp) == nil {
		return nil
	}

	target := *pp
	switch {
	case (*pp).Left != nil:
		pp = &(*pp).Left
		for (*pp).Right != nil {
			pp = &(*pp).Right
		}
		target.Val = (*pp).Val
		*pp = (*pp).Left

	case (*pp).Right != nil:
		pp = &(*pp).Right
		for (*pp).Left != nil {
			pp = &(*pp).Left
		}
		target.Val = (*pp).Val
		*pp = (*pp).Right

	default:
		*pp = nil
	}
}

func initLink(node *TreeLinkedNode) {
	node.Prev = node
	node.Next = node
}

func removeLink(node *TreeLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func insertLink(dummy, node *TreeLinkedNode) {
	p := dummy.Next
	for p != dummy && p.Val > node.Val {
		p = p.Next
	}

	// prev->p->next -> prev->p->node->next
	node.Prev = p
	p.Next.Prev = node
	node.Next = p.Next
	p.Next = node
}
