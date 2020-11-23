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

// BST with Double Linked from leetcode
type TreeLinkedNode struct {
	Val   int
	Left  *TreeLinkedNode
	Right *TreeLinkedNode
	Prev  *TreeLinkedNode
	Next  *TreeLinkedNode
}

func (root *TreeLinkedNode) kthSmallest(k int) int {
	kth := &TreeLinkedNode{}
	for p := root; p != nil; p = p.Next {
		if k -= 1; k == 0 {
			kth = p
			break
		}
	}
	return kth.Val
}
