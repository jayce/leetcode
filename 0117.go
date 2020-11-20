/*
 * Creation in 2020-11-20 by Chuan Liu (Jayce)
 */

package leetcode

// connect First idea with H space
func connect(root *Node) *Node {
	levels := []*Node{}                 // space: O(H)
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

// connect First idea without extra space
// but it's very slowly.
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	dfs(root.Left, root.Right)
	return root
}

func dfs(L, R *Node) {
	if L == nil || R == nil {
		connect(L)
		connect(R)
		return
	}

	L.Next = R
	dfs(R.Left, R.Right)
	dfs(L.Right, R.Right)
	dfs(L.Left, R.Right)

	dfs(L.Right, R.Left)
	dfs(L.Left, R.Left)

	dfs(L.Left, L.Right)
	return
}

// perfect BFS queue
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	dummy, queue := &Node{}, root // init: queue
	for queue != nil {            // time: O(H)
		for p := dummy; queue != nil; queue = queue.Next { // time: O(W)
			if queue.Left != nil {
				p.Next, p = queue.Left, queue.Left
			}
			if queue.Right != nil {
				p.Next, p = queue.Right, queue.Right
			}
		}
		queue = dummy.Next // next level
		dummy.Next = nil   // clean
	}
	return root
}
