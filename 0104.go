/*
 * Creation in 2020-11-10 by Chuan Liu (Jayce)
 */

package leetcode

// maxDepth First idea
func maxDepth(root *TreeNode) int { // time: O(N)
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Second idea
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root} // space: O(H)
	nodes := []*TreeNode{}     // space: O(H)
	for len(queue) > 0 {       // time: O(N)
		depth++
		nodes = nodes[0:0]
		for _, node := range queue {
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		queue, nodes = nodes, queue
	}
	return depth
}
