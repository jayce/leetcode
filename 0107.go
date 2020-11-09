/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

// levelOrderBottom First idea
func levelOrderBottom(root *TreeNode) [][]int {
	values := make([][]int, 0)
	levelOrderBottomHelp(root, 0, &values) // time: O(N)
	return values
}

func levelOrderBottomHelp(root *TreeNode, depth int, values *[][]int) {
	if root == nil {
		return
	}
	index := len(*values) - depth - 1
	if index < 0 { // add new the array of emtpy
		*values = append([][]int{{}}, (*values)...) // malloc: H times
		index = len(*values) - depth - 1
	}
	(*values)[index] = append((*values)[index], root.Val)
	levelOrderBottomHelp(root.Left, depth+1, values)
	levelOrderBottomHelp(root.Right, depth+1, values)
}

// Second idea optimize malloc
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	values := make([][]int, 0)
	aqueue := []*TreeNode{root} // space: O(width)
	bqueue := []*TreeNode{}     // space: O(width)
	for len(aqueue) > 0 {       // time: O(N)
		vs := make([]int, 0)
		bqueue = bqueue[0:0]
		for _, node := range aqueue {
			vs = append(vs, node.Val)
			if node.Left != nil {
				bqueue = append(bqueue, node.Left)
			}
			if node.Right != nil {
				bqueue = append(bqueue, node.Right)
			}
		}
		values = append(values, vs)
		aqueue, bqueue = bqueue, aqueue
	}
	return reverse(values) // time: O(H/2)
}

func reverse(values [][]int) [][]int {
	for s, e := 0, len(values)-1; s < e; s, e = s+1, e-1 {
		values[s], values[e] = values[e], values[s]
	}
	return values
}
