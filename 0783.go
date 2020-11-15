/*
 * Creation in 2020-11-15 by Chuan Liu (Jayce)
 */

package leetcode

// getMinimumDifference First idea
func getMinimumDifference(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	data := make([]int, 0)              // space: O(N)
	for len(stack) > 0 || root != nil { // time: O(N)
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		data = append(data, root.Val)
		root = root.Right
	}

	minimum := int(^uint32(0) >> 1)
	for i := 1; i < len(data); i++ { // time: O(N)
		if data[i]-data[i-1] < minimum {
			minimum = data[i] - data[i-1]
		}
	}
	return minimum
}
