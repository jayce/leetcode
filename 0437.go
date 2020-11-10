/*
 * Creation in 2020-11-10 by Chuan Liu (Jayce)
 */

package leetcode

// pathSum First idea
func pathSum(root *TreeNode, sum int) int { // time: O(N*N)
	if root == nil {
		return 0
	}
	return dfs(root, sum) + // time: O(N)
		pathSum(root.Left, sum) + // time: O(N-1)
		pathSum(root.Right, sum) // time: O(N-1)
}

func dfs(root *TreeNode, sum int) int { // time: O(N)
	pathCount := 0
	if root == nil {
		return pathCount
	}
	if root.Val == sum {
		pathCount++
	}
	return pathCount +
		dfs(root.Left, sum-root.Val) +
		dfs(root.Right, sum-root.Val)
}

// prefix sum of the path
func pathSum(root *TreeNode, sum int) int {
	stack := []*TreeNode{}         // space: O(H)
	prefixSumHash := map[int]int{} // space: O(N)
	prefixSumHash[0] = 1           // only zero in tree or prefixSum is zero

	prefixSum := 0
	pathCount := 0
	last := root

	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			prefixSum += root.Val
			pathCount += prefixSumHash[prefixSum-sum]
			prefixSumHash[prefixSum]++

			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		if root.Right != nil && root.Right != last {
			root = root.Right
			continue
		}

		// recover
		prefixSumHash[prefixSum]--
		prefixSum -= root.Val

		stack = stack[:len(stack)-1]
		last = root
		root = nil
	}

	return pathCount
}

// recursive
func pathSum(root *TreeNode, sum int) int {
	prefixSumHash := map[int]int{} // space: O(N)
	prefixSumHash[0] = 1           // only zero in tree or prefixSum is zero

	var dfsPrefixSum func(root *TreeNode, sum, prefixSum int) int
	dfsPrefixSum = func(root *TreeNode, prefixSum int) int {
		if root == nil {
			return 0
		}
		prefixSum += root.Val
		pathCount := prefixSumHash[prefixSum-sum]
		prefixSumHash[prefixSum]++
		pathCount += dfsPrefixSum(root.Left, prefixSum)
		pathCount += dfsPrefixSum(root.Right, prefixSum)
		prefixSumHash[prefixSum]-- // recover
		return pathCount
	}
	return dfsPrefixSum(root, 0) // time: O(N)
}
