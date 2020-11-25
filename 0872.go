/*
 * Creation in 2020-11-25 by Chuan Liu (Jayce)
 */

package leetcode

// leafSimilar First idea
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafValues := make([]int, 0) // space: O(M)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			leafValues = append(leafValues, root.Val)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root1) // time: O(N)
	dfs(root2) // time: O(N)

	if len(leafValues)%2 == 1 { // odd
		return false
	}

	for s, n := 0, len(leafValues)/2; s < n; s++ { // time: O(M/2)
		if leafValues[s] != leafValues[s+n] {
			return false
		}
	}
	return true
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	stack1 := append([]*TreeNode{}, root1)   // space: O(H)
	stack2 := append([]*TreeNode{}, root2)   // space: O(H)
	for len(stack1) > 0 && len(stack2) > 0 { // time: O(N)
		root1 = bfs(&stack1)
		root2 = bfs(&stack2)
		if root1.Val != root.Val {
			return false
		}
	}
	return len(stack1) == len(stack2)
}

func bfs(stack *[]*TreeNode) *TreeNode {
	for len(*stack) > 0 {
		node := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		if node.Left != nil {
			*stack = append(*stack, node.Left)
		}
		if node.Right != nil {
			*stack = append(*stack, node.Right)
		}
		if node.Left == nil && node.Right == nil {
			return node
		}
	}
	return nil
}
