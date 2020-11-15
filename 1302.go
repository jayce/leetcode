/*
 * Creation in 2020-11-13 by Chuan Liu (Jayce)
 */

package leetcode

// deepestLeavesSum First idea
func deepestLeavesSum(root *TreeNode) int {
	var dfs func(root *TreeNode, depth int)
	sums := make([]int, 0)                  // space: O(H)
	dfs = func(root *TreeNode, depth int) { // time: O(N)
		if root == nil {
			return
		}
		if len(sums) == depth {
			sums = append(sums, 0)
		}
		sums[depth] += root.Val
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return sums[len(sums)-1]
}

// optimized for space
func deepestLeavesSum(root *TreeNode) int {
	sum, maxDepth := 0, 0
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) { // time: O(N)
		switch {
		case root == nil:
			return
		case depth == maxDepth:
			sum += root.Val
		case depth > maxDepth:
			sum, maxDepth = root.Val, depth
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return sum
}

// Third idea
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	aqueue := []*TreeNode{root} // space: O(W)
	bqueue := []*TreeNode{}     // space: O(W)
	for len(aqueue) > 0 {       // time: O(N)
		sum = 0
		bqueue = bqueue[:0]
		for _, node := range aqueue {
			sum += node.Val
			if node.Left != nil {
				bqueue = append(bqueue, node.Left)
			}
			if node.Right != nil {
				bqueue = append(bqueue, node.Right)
			}
		}
		aqueue, bqueue = bqueue, aqueue
	}
	return sum
}
