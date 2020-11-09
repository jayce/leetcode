/*
 * Creation in 2020-11-09 by Chuan Liu (Jayce)
 */

package leetcode

// postorder First idea but will change the orignal data
func postorder(root *Node) []int {
	values := []int{}
	stack := []*Node{} // space: O(H)
	for root != nil {  // time: O(N)
		for len(root.Children) > 0 {
			stack = append(stack, root) // push root
			root = root.Children[0]
		}

		values = append(values, root.Val) // no children
		if len(stack) == 0 {              // no root
			break
		}

		root = stack[len(stack)-1]   // parent
		stack = stack[:len(stack)-1] // pop

		root.Children = root.Children[1:] // next
	}
	return values
}

// Second idea recursive
func postorder(root *Node) []int {
	return postorderHelp(root, &[]int{})
}

func postorderHelp(root *Node, nodes *[]int) []int {
	if root == nil {
		return *nodes
	}
	for _, node := range root.Children {
		postorderHelp(node, nodes)
	}
	*nodes = append(*nodes, root.Val)
	return *nodes
}

// reverse preorder
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	values := []int{}
	stack := []*Node{root} // space: O(H)
	for len(stack) > 0 {   // time: O(N)
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		values = append(values, root.Val)
		for i, size := 0, len(root.Children); i < size; i++ { // time: O(N)
			stack = append(stack, root.Children[i])
		}
	}
	return reverse(values) // time: O(N/2)
}

func reverse(values []int) []int {
	for s, e := 0, len(values)-1; s < e; s, e = s+1, e-1 {
		values[s], values[e] = values[e], values[s]
	}
	return values
}
