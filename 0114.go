/*
 * Creation in 2020-11-16 by Chuan Liu (Jayce)
 */

package leetcode

// flatten First idea
func flatten(root *TreeNode) { // time: O(N)
	if root == nil {
		return nil
	}
	flatten(root.Left)
	flatten(root.Right)
	root.Right = mergeTreeToRight(root.Left, root.Right) // time: O(N)
	root.Left = nil
}

func mergeTreeToRight(left, right *TreeNode) *TreeNode {
	if left == nil {
		return right
	}
	dummy := left
	for left != nil && left.Right != nil {
		left = left.Right
	}
	left.Right = right
	return dummy
}

// Second idea preorder, top-bottom
func flatten(root *TreeNode) {
	lastNode := &TreeNode{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) { // time: O(N)
		if root == nil {
			return
		}
		lastNode.Right = root
		lastNode.Left = nil
		lastNode = root
		right := root.Right // why?
		dfs(root.Left)
		dfs(right)
	}
	dfs(root)
}

// Third idea postorder, bottom-top
func flatten(root *TreeNode) {
	lastNode := (*TreeNode)(nil)
	stack := []*TreeNode{}              // space: O(H)
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}

		root = stack[len(stack)-1]
		if root.Left != nil && root.Left != lastNode {
			root = root.Left
			continue
		}
		stack = stack[:len(stack)-1]
		root.Right = lastNode
		lastNode = root
		lastNode.Left = nil
		root = nil
	}
}

// Fourth idea from No.206 ReverseListIII
func flatten(root *TreeNode) {
	var postorder func(root, tail *TreeNode) *TreeNode
	postorder = func(root, tail *TreeNode) *TreeNode { // time: O(N)
		if root == nil {
			return tail
		}
		root.Right = postorder(root.Right, tail)
		root.Right = postorder(root.Left, root.Right)
		root.Left = nil
		return root
	}
	postorder(root, nil)
}
