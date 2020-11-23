/*
 * Creation in 2020-11-23 by Chuan Liu (Jayce)
 */

package leetcode

// removeLeafNodes First idea
func removeLeafNodes(root *TreeNode, target int) *TreeNode { // time: O(N)
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Val == target && root.Left == root.Right {
		return nil
	}
	return root
}

// Second idea, postorder bottom-top
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	stack := []*TreeNode{} // time: O(H)
	lastNode := (*TreeNode)(nil)
	dummy := &TreeNode{Left: root}
	root = dummy

	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]

		if lastNode != nil && lastNode.Val == target &&
			lastNode.Left == lastNode.Right {
			if root.Left == lastNode {
				root.Left = nil
			}
			if root.Right == lastNode {
				root.Right = nil
			}
		}

		if root.Right != nil && root.Right != lastNode {
			root = root.Right
			continue
		}

		stack = stack[:len(stack)-1]
		lastNode = root
		root = nil
	}

	return dummy.Left // or lastNode.Left
}
