/*
 * Creation in 2020-11-23 by Chuan Liu (Jayce)
 */

package leetcode

// longestUnivaluePath First idea, top-bottom
func longestUnivaluePath(root *TreeNode) int {
	maxPathDepth := 0
	var dfs func(root, p *TreeNode, depth int) int
	dfs = func(root, p *TreeNode, depth int) int { // time: O(N)
		if root == nil {
			return depth
		}
		if p != nil && root.Val == p.Val {
			L := dfs(root.Left, root, depth+1)
			R := dfs(root.Right, root, depth+1)
			maxPathDepth = max(maxPathDepth, L+R-(depth+1)*2) // minus next level value
			return max(L, R)
		}
		L := dfs(root.Left, root, 0)
		R := dfs(root.Right, root, 0)
		maxPathDepth = max(maxPathDepth, L+R)
		return depth
	}
	dfs(root, nil, 0)
	return maxPathDepth
}

// Second idea more clear, bottom-top
func longestUnivaluePath(root *TreeNode) int {
	maxPath := 0
	var depth func(root *TreeNode, val int) int
	depth = func(root *TreeNode, val int) int { // time: O(N)
		if root == nil {
			return 0
		}
		L := depth(root.Left, root.Val)
		R := depth(root.Right, root.Val)
		maxPath = max(maxPath, L+R)
		if root.Val == val {
			return max(L, R) + 1
		}
		return 0
	}
	depth(root, root.Val)
	return maxPath
}
