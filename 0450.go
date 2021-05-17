/*
 * Creation in 2021-05-06 by Chuan Liu (Jayce)
 */

package leetcode

// First idea
func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
		return root
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
		return root

	// root.val == key
	case root.Left == nil:
		return root.Right
	case root.Right == nil:
		return root.Left
	}

	minimum := (*TreeNode)(nil)
	root.Right, minimum = deleteMinimum(root.Right)
	minimum.Right = root.Right
	minimum.Left = root.Left
	return minimum
}

func deleteMinimum(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}

	if root.Left == nil {
		return root.Right, root
	}

	minimum := (*TreeNode)(nil)
	root.Left, minimum = deleteMinimum(root.Left)
	return root, minimum
}

// lterative, double pointer, inspired by other issue
func deleteNode(root *TreeNode, key int) *TreeNode {
	pp := &root
	for *pp != nil && (*pp).Val != key {
		switch {
		case (*pp).Val > key:
			pp = &(*pp).Left
		case (*pp).Val < key:
			pp = &(*pp).Right
		}
	}

	if (*pp) == nil {
		return root
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

	return root
}
