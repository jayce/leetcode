/*
 * Creation in 2020-11-24 by Chuan Liu (Jayce)
 */

package leetcode

// increasingBST First idea Left first, top-bottom
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root = dfs(root)
	newroot := root.Left
	root.Left = nil
	return newroot
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	L := dfs(root.Left)
	if L == nil {
		L = root
		L.Left = root
	} else {
		newroot := L.Left // pop newroot
		L.Left = nil      // remove newroot

		L.Right = root
		L = L.Right

		L.Left = newroot // save newroot
	}

	R := dfs(root.Right)
	if R != nil {
		L.Right = R.Left
		R.Left = L.Left // save newroot
		L.Left = nil    // remove newroot
		return R
	}

	return L
}

// Second idea inorder but from Right to Left, bottom-top
func increasingBST(root *TreeNode) *TreeNode {
	lastNode := (*TreeNode)(nil)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) { // time: O(N)
		if root == nil {
			return nil
		}
		inorder(root.Right)
		root.Right = lastNode
		lastNode = root
		root = root.Left
		lastNode.Left = nil
		inorder(root)
	}
	inorder(root)
	return lastNode
}

// Third idea more clear than First idea, top-bottom
func increasingBST(root *TreeNode) *TreeNode {
	newRoot := (*TreeNode)(nil)
	lastNode := &newRoot
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) { // time: O(N)
		if root == nil {
			return
		}
		inorder(root.Left)
		if *lastNode != nil {
			(*lastNode).Right = root
		}
		root.Left = nil
		*lastNode = root
		lastNode = &root
		inorder(root.Right)
	}
	inorder(root)
	return newRoot
}

// From No.114 same as Second idea but non-recursive
func increasingBST(root *TreeNode) *TreeNode {
	lastNode := (*TreeNode)(nil)
	stack := []*TreeNode{}              // space: O(H)
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root.Right = lastNode
		lastNode = root
		root = root.Left
		lastNode.Left = nil
	}
	return lastNode
}

// bottom-top, insert from list's head always
func increasingBST(root *TreeNode) *TreeNode {
	var reverseInorder func(root, head *TreeNode) *TreeNode
	reverseInorder = func(root, head *TreeNode) *TreeNode { // time: O(N)
		if root == nil {
			return head
		}
		root.Right = reverseInorder(root.Right, head)
		newHead := reverseInorder(root.Left, root)
		root.Left = nil
		return newHead
	}
	return reverseInorder(root, nil)
}
