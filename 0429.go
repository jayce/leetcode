/*
 * Creation in 2020-11-13 by Chuan Liu (Jayce)
 */

package leetcode

// levelOrder First idea
func levelOrder(root *Node) [][]int {
	rows := [][]int{}
	var dfs func(root *Node, depth int)
	dfs = func(root *Node, depth int) {
		if root == nil {
			return
		}
		if len(rows) == depth {
			rows = append(rows, []int{})
		}
		rows[depth] = append(rows[depth], root.Val)
		for _, node := range root.Children {
			dfs(node, depth+1)
		}
	}

	dfs(root, 0)
	return rows
}

// Second idea
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	rows := [][]int{}
	aqueue := []*Node{root}
	bqueue := []*Node{}
	depth := 0

	for len(aqueue) > 0 { // time: O(N)
		if len(rows) == depth {
			rows = append(rows, []int{})
		}
		bqueue = bqueue[:0]
		for _, node := range aqueue {
			rows[depth] = append(rows[depth], node.Val)
			if len(node.Children) > 0 {
				bqueue = append(bqueue, node.Children...) // copy memory: O(N)
			}
		}
		aqueue, bqueue = bqueue, aqueue
		depth++
	}

	return rows
}
