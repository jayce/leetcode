/*
 * Creation in 2020-11-13 by Chuan Liu (Jayce)
 */

package leetcode

// lowestCommonAncestor First idea
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPaths := dfs(root, p)             // time: O(N) space: O(H)
	qPaths := dfs(root, q)             // time: O(N) space: O(H)
	for i := 0; i < len(pPaths); i++ { // time: O(H*H)
		for j := 0; j < len(qPaths); j++ {
			if pPaths[i] == qPaths[j] {
				return qPaths[j]
			}
		}
	}
	return nil
}

func dfs(root, p *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return []*TreeNode{root}
	}
	paths := dfs(root.Left, p)
	if paths != nil && len(paths) > 0 {
		paths = append(paths, root)
		return paths
	}

	paths = dfs(root.Right, p)
	if paths != nil && len(paths) > 0 {
		paths = append(paths, root)
		return paths
	}
	return nil
}

// Second idea time: O(N)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	L := lowestCommonAncestor(root.Left, p, q)
	R := lowestCommonAncestor(root.Right, p, q)
	if L != nil && R != nil {
		return root
	}
	if L != nil {
		return L
	}
	return R
}
