/*
 * Creation in 2020-11-10 by Chuan Liu (Jayce)
 */

package leetcode

// minDepth First idea
func minDepth(root *TreeNode) int { // time: O(N)
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Second idea
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	aqueue := []*TreeNode{root} // space: O(minDepth)
	bqueue := []*TreeNode{}     // space: O(minDepth)
	for len(aqueue) > 0 {       // time: O(N)
		depth++
		bqueue = bqueue[:0]
		for _, node := range aqueue {
			if node.Left == node.Right { // both nil
				return depth
			}
			if node.Left != nil {
				bqueue = append(bqueue, node.Left)
			}
			if node.Right != nil {
				bqueue = append(bqueue, node.Right)
			}
		}
		aqueue, bqueue = bqueue, aqueue
	}
	return depth
}
