/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

// levelOrder First idea
func levelOrder(root *TreeNode) [][]int {
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
	return values
}

// Second idea recursive
func levelOrder(root *TreeNode) [][]int { // time: O(N)
	values := make([][]int, 0)
	levelOrderBottomHelp(root, 0, &values)
	return values
}

func levelOrderBottomHelp(root *TreeNode, depth int, values *[][]int) {
	if root == nil {
		return
	}
	if depth > len(*values)-1 {
		*values = append(*values, []int{})
	}
	(*values)[depth] = append((*values)[depth], root.Val)
	levelOrderBottomHelp(root.Left, depth+1, values)
	levelOrderBottomHelp(root.Right, depth+1, values)
}
