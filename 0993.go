/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

// isCousins First idea
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	aqueue, bqueue := []*TreeNode{root}, []*TreeNode{} // space: O(2 * width)
	xnode, ynode := root, root                         // just for init

	for len(aqueue) > 0 { // time: O(N)
		xnode, ynode = nil, nil
		bqueue = bqueue[:0]

		for _, node := range aqueue {
			compare(&xnode, node, x)
			compare(&ynode, node, y)
			if xnode != nil && ynode != nil {
				return xnode != ynode
			}

			if node.Left != nil {
				bqueue = append(bqueue, node.Left)
			}

			if node.Right != nil {
				bqueue = append(bqueue, node.Right)
			}
		}

		if xnode != nil || ynode != nil { // found x or y
			return xnode != nil && ynode != nil && xnode != ynode
		}

		aqueue, bqueue = bqueue, aqueue
	}

	return false
}

func compare(node **TreeNode, root *TreeNode, n int) {
	if (root.Left != nil && root.Left.Val == n) ||
		(root.Right != nil && root.Right.Val == n) {
		*node = root
	}
}

// Second idea
func isCousins(root *TreeNode, x int, y int) bool {
	xroot, xdepth := dfs(root, root.Val, x) // time: O(N)
	yroot, ydepth := dfs(root, root.Val, y) // time: O(N)
	return xroot != yroot && xdepth == ydepth
}

func dfs(root *TreeNode, val, n int) (int, int) {
	if root == nil {
		return val, -1
	}
	if root.Val == n {
		return val, 1
	}
	lv, ld := dfs(root.Left, root.Val, n)
	if ld > 0 {
		return lv, ld + 1
	}
	rv, rd := dfs(root.Right, root.Val, n)
	if rd > 0 {
		return rv, rd + 1
	}
	return rv, rd
}
