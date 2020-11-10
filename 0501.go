/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

// findMode First idea
func findMode(root *TreeNode) []int { // time: O(2N)
	modes := make([]int, 0)
	maxRepeat, repeat := 0, 0
	last := root

	dfs(root, func(root *TreeNode) { // time: O(N)
		if last != nil && last.Val != root.Val {
			repeat = 0
		}
		last = root
		repeat++
		maxRepeat = max(maxRepeat, repeat)
	})

	repeat, root = 0, nil
	dfs(root, func(root *TreeNode) { // time: O(N)
		if last != nil && last.Val != root.Val {
			repeat = 0
		}
		last = root
		repeat++
		if maxRepeat == repeat {
			modes = append(modes, root.Val)
		}
	})
	return modes
}

func dfs(root *TreeNode, f func(root *TreeNode)) {
	if root == nil {
		return
	}
	dfs(root.Left, f)
	f(root)
	dfs(root.Right, f)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Second idea optimize
func findMode(root *TreeNode) []int {
	modes := make([]int, 0)
	maxRepeat, repeat := 0, 0
	last := root

	dfs(root, func(root *TreeNode) { // time: O(N)
		if last != nil && root.Val != last.Val {
			repeat = 0
		}
		repeat++
		last = root
		if maxRepeat < repeat {
			maxRepeat = repeat
			modes = modes[:0]
		}
		if maxRepeat == repeat {
			modes = append(modes, root.Val)
		}
	})
	return modes
}

// Third idea map
func findMode(root *TreeNode) []int {
	ranks := map[int]int{}
	modes := make([]int, 0, 1024)
	maxRepeat := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		// inorder traversal
		ranks[root.Val]++
		dfs(root.Right)
	}
	dfs(root)

	for k, v := range ranks {
		if maxRepeat == v {
			modes = append(modes, k)
		} else if maxRepeat < v {
			maxRepeat = v
			modes = modes[:0]
			modes = append(modes, k)
		}
	}
	return modes
}
