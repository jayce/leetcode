/*
 * Creation in 2020-11-16 by Chuan Liu (Jayce)
 */

package leetcode

// addOneRow First idea
func addOneRow(root *TreeNode, v int, d int) *TreeNode { // time: O(N)
	if root == nil {
		return nil
	}
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	if d == 2 {
		root.Left = &TreeNode{Val: v, Left: root.Left}
		root.Right = &TreeNode{Val: v, Right: root.Right}
		return root
	}

	root.Left = addOneRow(root.Left, v, d-1)
	root.Right = addOneRow(root.Right, v, d-1)
	return root
}

// Second idea
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}

	aqueue := make([]*TreeNode, 0) // space: O(H)
	bqueue := make([]*TreeNode, 0) // space: O(H)
	aqueue = append(aqueue, root)
	for len(aqueue) > 0 && d > 2 { // time: O(N)
		d--
		bqueue = bqueue[0:0]
		for _, node := range aqueue {
			if node.Left != nil {
				bqueue = append(bqueue, node.Left)
			}
			if node.Right != nil {
				bqueue = append(bqueue, node.Right)
			}
		}
		aqueue, bqueue = bqueue, aqueue
	}

	for _, node := range aqueue { // time: O(W)
		node.Left = &TreeNode{Val: v, Left: node.Left}
		node.Right = &TreeNode{Val: v, Right: node.Right}
	}

	return root
}
