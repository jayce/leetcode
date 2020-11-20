/*
 * Creation in 2020-11-19 by Chuan Liu (Jayce)
 */

package leetcode

// connect First idea
func connect(root *Node) *Node {
	levels := make([]*Node, 0)          // space: O(H)
	var dfs func(root *Node, depth int) // time: O(N)
	dfs = func(root *Node, depth int) {
		if root == nil {
			return
		}

		if depth == len(levels) { // why put it there?
			levels = append(levels, root)
		} else {
			levels[depth].Next = root
			levels[depth] = root
		}

		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return root
}

// Second idea
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	dfs(root.Left, root.Right) // time: O(N)
	return root
}
func dfs(L, R *Node) {
	if L == nil || R == nil {
		return
	}
	L.Next = R           // same depth
	dfs(L.Left, L.Right) // same root
	dfs(R.Left, R.Right)
	dfs(L.Right, R.Left) // cross
	return
}

func connect(root *Node) *Node { // time: O(N)
	if root == nil {
		return root
	}
	L, R := root.Left, root.Right
	for L != nil {
		L.Next = R
		L = L.Right
		R = R.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
