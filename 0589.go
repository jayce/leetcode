/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

// preorder First idea
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	values := []int{}
	stack := []*Node{root} // space: O(H)
	for len(stack) > 0 {   // time: O(N)
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		values = append(values, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- { // time: O(N)
			stack = append(stack, root.Children[i])
		}
	}
	return values
}

// Second idea recursive
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	values := []int{root.Val}
	for _, node := range root.Children {
		values = append(values, preorder(node)...)
	}
	return values
}
