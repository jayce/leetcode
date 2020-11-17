/*
 * Creation in 2020-11-17 by Chuan Liu (Jayce)
 */

package leetcode

// insertIntoMaxTree First idea
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || val >= root.Val {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}

	// val < root.Val, only concern the Right
	root.Right = insertIntoMaxTree(root.Right, val) // time: O(logN) or O(N)
	return root
}

// Second idea just another version From the First
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	dummy := &TreeNode{Right: root}
	stack := []*TreeNode{dummy}                              // space: O(logN) or O(N)
	for ; root != nil && val < root.Val; root = root.Right { // time: O(logN) or O(N)
		stack = append(stack, root)
	}
	stack[len(stack)-1].Right = &TreeNode{
		Val:  val,
		Left: root,
	}

	return dummy.Right
}
